package dataapi

import (
	"context"
)

// NukeDatabase is a data function will delete all the Data API kinds in the template-infrastructure namespace
// Usage: [NukeDatabase()]
// The function will delete most entities from the namespace template-infrastructure
func (d DataAPIService) NukeDatabase() interface{} {

	// Delete DataAPI kinds from the automated-tests namespace
	err := d.Nuke.DeleteAutomatedTests(context.Background())
	if err != nil {
		return err
	}

	// Return success
	return nil

}