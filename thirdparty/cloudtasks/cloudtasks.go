package cloudtasks

import (
	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"context"
	"github.com/Celbux/template-infrastructure/business/i"
)

// NewClient creates a GCP CloudTasks client
func NewClient(ctx context.Context, log i.Logger) (*cloudtasks.Client, error) {

	ctClient, err := cloudtasks.NewClient(ctx)

	if err != nil {
		log.Printf("Cloud Tasks Failed to Initialize, %s", err)
		return nil, err
	}
	log.Println("Cloud Tasks initialized")
	return ctClient, err
}
