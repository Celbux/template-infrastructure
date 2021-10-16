package dataapi

import (
	"context"
	"github.com/Celbux/dataapi/business/dataapi"
	"github.com/Celbux/template-infrastructure/business/i"
)

// DataAPIService encapsulates all dependencies required by the DataAPI.
// This service is used to run data driven functionality at run time
type DataAPIService struct {
	CoreDataAPI dataapi.DataAPIService
	Nuke 		Nuke
	Log         i.Logger
}

// Nuke is responsible for deleting data from a namespace
type Nuke interface {
	DeleteAutomatedTests(ctx context.Context) error
}

type EvalCache map[string]interface{}
