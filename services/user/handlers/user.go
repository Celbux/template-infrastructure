package handlers

import (
	"context"
	"github.com/Celbux/template-infrastructure/business/user"
	"github.com/Celbux/template-infrastructure/foundation/web"
	"net/http"
)

type UserHandlers struct {
	Service user.UserService
}

func (u UserHandlers) createUser(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
) error {

	// Get first and last name from the post request
	type request struct {
		FirstName string `json:"FirstName"`
		LastName string `json:"LastName"`
		Data string `json:"Data"`
	}
	req := request{}
	err := web.Decode(r, &req)
	if err != nil {
		return web.NewError(err)
	}

	// Create the user
	err = u.Service.CreateUser(ctx, req.FirstName, req.LastName, req.Data)
	if err != nil {
		return err
	}

	// Return success
	return web.Respond(ctx, w, nil, http.StatusNoContent)

}

func (u UserHandlers) getUser(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
) error {

	// Get ID of the user we want to retrieve
	type request struct {
		ID string `json:"Id"`
	}
	req := request{}
	err := web.Decode(r, &req)
	if err != nil {
		return web.NewError(err)
	}

	// Get the user
	entity, err := u.Service.GetUser(ctx, req.ID)
	if err != nil {
		return err
	}

	// Return success
	return web.Respond(ctx, w, entity, http.StatusOK)

}