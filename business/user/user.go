package user

import (
	"context"
	"github.com/Celbux/template-infrastructure/foundation/web"
	"strings"
)

// CreateUser will Capitalize both users names and write the record
func (u UserService) CreateUser(ctx context.Context, firstName string, lastName string, data string) error {

	// Capitalize the names
	firstName = strings.ToUpper(firstName)
	lastName = strings.ToUpper(lastName)

	// Write to Datastore
	err := u.CRUD.CreateUser(ctx, firstName, lastName, data)
	if err != nil {
		return web.NewError(err)
	}

	// Return success
	return nil

}

// GetUser will retrieve the user
func (u UserService) GetUser(ctx context.Context, id string) (interface{}, error) {

	// Write to Datastore
	user, err := u.CRUD.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// Return success
	return user, nil

}
