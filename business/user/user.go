package user

import (
	"context"
	"github.com/Celbux/template-infrastructure/business/i"
)

// Service contains all methods that can be performed on a user
type Service struct {
	Log   i.Logger
	Store Store
}

// Store encapsulates third party dependencies
type Store interface {
	CreateUser(ctx context.Context, firstName string, lastName string, data string) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}