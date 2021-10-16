package datastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/Celbux/template-infrastructure/foundation/web"
)

type UserStore struct {
	DB *datastore.Client
}

func (u UserStore) CreateUser(ctx context.Context, firstName string, lastName string, data string) error {

	// Create the entity and the key
	entity := struct {
		FirstName string
		LastName  string
		Data      string
	}{
		FirstName: firstName,
		LastName:  lastName,
		Data:      data,
	}
	namespace := ctx.Value("namespace").(string)
	key := &datastore.Key{
		Kind: "User",
		Name: firstName + lastName,
		Namespace: namespace,
	}

	// Insert the entity into datastore
	_, err := u.DB.Put(ctx, key, &entity)
	if err != nil {
		return web.NewError(err)
	}

	// Return success
	return nil

}

func (u UserStore) GetUser(ctx context.Context, id string) (interface{}, error) {

	// Retrieve entity as property load saver
	namespace := ctx.Value("namespace").(string)
	key := &datastore.Key{
		Kind: "User",
		Name: id,
		Namespace: namespace,
	}
	var propertyList datastore.PropertyList
	err := u.DB.Get(ctx, key, &propertyList)
	if err != nil {
		return nil, web.NewError(err)
	}

	// Retrieve properties from the property load saver and return as interface
	response := make(map[string]interface{})
	for _, property := range propertyList {
		response[property.Name] = property.Value
	}

	// Return success
	return response, nil

}
