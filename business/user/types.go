package user

import (
	"context"
	"github.com/Celbux/template-infrastructure/business/i"
)

// UserService contains all methods that can be performed on a user
type UserService struct {
	CRUD UserStore
	Log i.Logger
}

// UserStore encapsulates third party dependencies
type UserStore interface {
	CreateUser(ctx context.Context, firstName string, lastName string, data string) error
	GetUser(ctx context.Context, id string) (interface{}, error)
}


