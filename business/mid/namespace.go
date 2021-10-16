package mid

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	"github.com/Celbux/template-infrastructure/business/i"
	"github.com/Celbux/template-infrastructure/foundation/web"
)

// Namespace will set the namespace on the context
func Namespace(log i.Logger) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

			// If the context is missing this value, request the service to be shutdown gracefully
			_, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			// Get namespace from url parameter
			namespace := web.GetParam(r, "ns")
			if namespace == "" || namespace == "default" {
				return &web.Error{Err: errors.New("namespace cannot be empty or default")}
			}

			// Add namespace to ctx
			ctx = context.WithValue(ctx, "namespace", namespace)

			// Call the next handler and set its return value in the err variable.
			return handler(ctx, w, r)
		}

		return h
	}

	return m
}