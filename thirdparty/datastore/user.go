package datastore

import (
	ds "cloud.google.com/go/datastore"
	"context"
	"github.com/Celbux/template-infrastructure/business/user"
	"github.com/Celbux/template-infrastructure/foundation/web"
)

type UserStore struct {
	DB *ds.Client
}

func (u UserStore) CreateUser(ctx context.Context, firstName string, lastName string, data string) (*user.User, error) {

	// Get namespace from ctx
	namespace := ctx.Value("namespace").(string)

	// Create User
	entity := &user.User{
		FirstName: firstName,
		LastName:  lastName,
		Data:      data,
	}

	// Create name key
	key := ds.NameKey("User", firstName + lastName, nil)
	key.Namespace = namespace

	// Insert the entity into datastore
	_, err := u.DB.Put(ctx, key, entity)
	if err != nil {
		return nil, web.NewError(err)
	}

	// Return user & success
	return entity, nil

}

func (u UserStore) GetUser(ctx context.Context, id string) (*user.User, error) {

	// Get namespace from ctx
	namespace := ctx.Value("namespace").(string)

	// Create name key
	key := ds.NameKey("User", id, nil)
	key.Namespace = namespace

	// Get entity via property load saver
	var propertyList ds.PropertyList
	err := u.DB.Get(ctx, key, &propertyList)
	if err != nil {
		return nil, web.NewError(err)
	}

	// Retrieve properties from the property load saver
	response := make(map[string]interface{})
	for _, property := range propertyList {
		response[property.Name] = property.Value
	}

	// Build User
	entity := &user.User{
		FirstName: response["FirstName"].(string),
		LastName:  response["LastName"].(string),
		Data:      response["Data"].(string),
	}

	// Return user & success
	return entity, nil

}
