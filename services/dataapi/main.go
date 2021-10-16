package main

import (
	"context"
	"fmt"
	coredataapi "github.com/Celbux/dataapi/business/dataapi"
	"github.com/Celbux/template-infrastructure/business/dataapi"
	"github.com/Celbux/template-infrastructure/services/dataapi/handlers"
	ds "github.com/Celbux/template-infrastructure/thirdparty/datastore"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
)

func main() {

	err := run(log.New(os.Stdout, "", 0))
	if err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}

}

func run(log *log.Logger) error {

	// Configuration uses github.com/ardanlabs/conf library
	// Your program configuration is attempted to be retrieved in the priority:
	// 1) Environment variable
	// 2) CMD flag
	// 3) Else the default value will be used
	defer log.Println("main: Completed")
	var cfg struct {
		Web struct {
			APIHost         string        `conf:"default:0.0.0.0:8082"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:0s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
		Datastore struct {
			ProjectID           string `conf:"default:dev8celbux"`
			Setting             string `conf:"default:LOCAL_WITH_CLOUD_DB,help:options - CLOUD | LOCAL_WITH_CLOUD_DB | LOCAL"`
			EmulatorHost        string `conf:"default:localhost:8080,noprint"`
			CredentialsFilePath string `conf:"default:./key.json"`
		}
	}
	namespace := "TEMPLATE_INFRASTRUCTURE_DATA_API"
	log.Printf("main: Starting : %v\n", namespace)
	if err := conf.Parse(os.Args[1:], namespace, &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage(namespace, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString(namespace, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}
	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main: Config :\n%v\n", out)
	dbSetting := ds.StringToSetting(log, cfg.Datastore.Setting)

	// Dependency Injection: Create our Services with their dependencies to
	// attach on for later access via receiver functions
	ctx := context.Background()
	log.Println("main: Initializing database support")
	dsClient, err := ds.NewClient(ctx, log, ds.Config{
		ProjectID:           cfg.Datastore.ProjectID,
		EmulatorHost:        cfg.Datastore.EmulatorHost,
		CredentialsFilePath: cfg.Datastore.CredentialsFilePath,
		Setting:             dbSetting,
	})
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}

	// Shut down all services graceful in a defer
	defer func() {
		log.Printf("main: Database Stopping : %s", cfg.Datastore.ProjectID)
		err := dsClient.Close()
		if err != nil {
			log.Printf("main: Stopping database ERROR: %s", err.Error())
		}
	}()

	dataAPI := handlers.DataAPIHandlers{
		Service: dataapi.DataAPIService{
			CoreDataAPI: coredataapi.DataAPIService{Log: log},
			Log:         log,
			Nuke:        ds.Nuke{
				Log: log,
				DB:  dsClient,
			},
		},
	}
	log.Println("main: Initializing API Handlers")

	// Make a channel to listen for an interrupt or terminate signal from the
	// OS. Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr: cfg.Web.APIHost,
		Handler: handlers.API(
			dataAPI,
			log,
			shutdown,
		),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this
	// error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(
			context.Background(),
			cfg.Web.ShutdownTimeout,
		)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := api.Shutdown(ctx); err != nil {
			_ = api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil

}
