package datastore

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"

	"github.com/Celbux/template-infrastructure/business/i"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

// Config is the required properties to use the database.
type Config struct {
	ProjectID           string
	EmulatorHost        string
	CredentialsFilePath string
	Setting             int
}

const (
	// Cloud means our app is running in GCP so we don't
	// have to pass credentials or set ENV variables
	Cloud = iota

	// LocalWithCloudDb means we are running against Cloud Datastore
	// from our local machine, so we need a credentials file
	LocalWithCloudDb = iota

	// Local means we are running against a Datastore Emulator,
	// so we don't need credentials
	Local = iota
)

// NewClient knows how to open a database connection based on the configuration.
// A lot of local-dev environment vs production environment logic gets configured here
func NewClient(
	ctx context.Context,
	log i.Logger,
	cfg Config,
) (*datastore.Client, error) {

	// When running in GCP we don't need any credentials or ENV variables set
	if cfg.Setting == Cloud {
		return cloudDatastore(ctx, cfg.ProjectID)
	}

	// When running locally we need to set some ENV variables depending on how
	// we want to run it. This is purely for ease of development
	localhostErr := prepEnvironment(
		cfg.Setting == Local,
		cfg.ProjectID,
		cfg.EmulatorHost,
		cfg.CredentialsFilePath,
		log,
	)

	if localhostErr != nil {
		return nil, localhostErr
	}

	// If we're using the local emulator we can return this option
	if cfg.Setting == Local {
		return datastoreEmulator(ctx, cfg.ProjectID)
	}

	if cfg.Setting == LocalWithCloudDb {
		return cloudDatastore(ctx, cfg.ProjectID)
	}

	return nil, errors.New("DB connection failed for unknown reason")

}

func cloudDatastore(
	ctx context.Context,
	projectID string,
) (*datastore.Client, error) {
	return datastore.NewClient(ctx, projectID)
}

func datastoreEmulator(
	ctx context.Context,
	projectID string,
) (*datastore.Client, error) {
	return datastore.NewClient(ctx, projectID, option.WithoutAuthentication())
}

// prepEnvironment sets or unsets env variables for datastore emulator
func prepEnvironment(
	useLocalEmulator bool,
	projectID,
	emulatorHost,
	credentialsFilePath string,
	log i.Logger,
) error {
	if useLocalEmulator {
		log.Println("Setting Datastore Emulator ENV Variables")

		_ = os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")
		_ = os.Setenv("DATASTORE_DATASET", projectID)
		_ = os.Setenv("DATASTORE_EMULATOR_HOST", emulatorHost)
		_ = os.Setenv("DATASTORE_EMULATOR_HOST_PATH", fmt.Sprintf("%s/datastore", emulatorHost))
		_ = os.Setenv("DATASTORE_PROJECT_ID", projectID)

		request, _ := http.NewRequest("GET", fmt.Sprintf("http://%s", emulatorHost), nil)
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil || resp.StatusCode != http.StatusOK {
			return errors.Errorf(`
			Datastore emulator not running at %s
			Try running "gcloud beta emulators datastore start"
			More information on datastore emulator: https://cloud.google.com/datastore/docs/tools/datastore-emulator
			`, emulatorHost)
		}
	} else {
		log.Println("Unsetting Datastore Emulator ENV Variables")
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentialsFilePath)
		_ = os.Unsetenv("DATASTORE_DATASET")
		_ = os.Unsetenv("DATASTORE_EMULATOR_HOST")
		_ = os.Unsetenv("DATASTORE_EMULATOR_HOST_PATH")
		_ = os.Unsetenv("DATASTORE_PROJECT_ID")
	}
	return nil
}

// StringToSetting validates the db settings configuration
// and converts it to const values
func StringToSetting(log i.Logger, setting string) int {

	if setting == "CLOUD" {
		return Cloud
	}

	if setting == "LOCAL_WITH_CLOUD_DB" {
		return LocalWithCloudDb
	}

	if setting == "LOCAL" {
		return Local
	}

	log.Println("")
	log.Println("")
	log.Println("Warning: Datastore.Setting config value defaulting to CLOUD for production safety.")
	log.Println("Datastore setting has to be either CLOUD, LOCAL_WITH_CLOUD_DB, or LOCAL")
	log.Println("")
	log.Println("")

	return Cloud
}
