package user

import (
	"context"
	"github.com/Celbux/template-infrastructure/foundation/web"
	"strings"
)

// CreateUser will Capitalize both users names and write the record
func (s Service) CreateUser(ctx context.Context, firstName string, lastName string, data string) error {

	// Capitalize the names
	firstName = strings.ToUpper(firstName)
	lastName = strings.ToUpper(lastName)

	// Write to Datastore
	_, err := s.Store.CreateUser(ctx, firstName, lastName, data)
	if err != nil {
		return web.NewError(err)
	}

	// Return success
	return nil

}

// GetUser will retrieve the user
func (s Service) GetUser(ctx context.Context, id string) (*User, error) {

	// Get user with username
	user, err := s.Store.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// Return success
	return user, nil

}
