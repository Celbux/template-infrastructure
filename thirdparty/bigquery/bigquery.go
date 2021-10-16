package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
)

// Config is the required properties to use the database.
type Config struct {
	ProjectID string
}

// NewClient initializes and returns a new BigQuery client, if the connection
// fails an error is returned
func NewClient(
	ctx context.Context,
	cfg Config,
) (*bigquery.Client, error) {

	bqClient, err := bigquery.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		return nil, err
	}

	return bqClient, err
}
