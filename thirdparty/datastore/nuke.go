package datastore

import (
	ds "cloud.google.com/go/datastore"
	"context"
	"github.com/Celbux/template-infrastructure/business/i"
	"github.com/Celbux/template-infrastructure/foundation/tools"
	"github.com/Celbux/template-infrastructure/thirdparty/console"
	"google.golang.org/api/iterator"
)

type Nuke struct {
	Log i.Logger
	DB *ds.Client
}

func (n Nuke) DeleteAutomatedTests(ctx context.Context) error {

	// Delete only the kinds in the template-infrastructure namespace
	envKinds := []string{"User"}

	// Track progress of deletion
	n.Log.Println("Deleting template-infrastructure Namespace...")
	progress := console.NewProgressBar()

	var keys []*ds.Key
	for i, kind := range envKinds {
		// Update progress bar
		progress.Update(i+1, len(envKinds))

		// Get all keys from each kind and build up an array to bulk delete
		query := ds.NewQuery(kind).KeysOnly().Namespace("template-infrastructure")
		it := n.DB.Run(ctx, query)
		for {
			var key ds.Key
			keyVal, err := it.Next(&key)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			keys = append(keys, keyVal)
		}

		// Delete all records in chunks of 500 or less
		for i := 0; i < len(keys); i += 500 {
			chunk := tools.Min(len(keys)-i, 500)
			err := n.DB.DeleteMulti(ctx, keys[i:i+chunk])
			if err != nil {
				return err
			}
		}
	}

	// Return success
	return nil

}
