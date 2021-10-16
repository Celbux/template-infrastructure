package handlers

import (
	ds "cloud.google.com/go/datastore"
	"github.com/Celbux/template-infrastructure/business/i"
	"github.com/Celbux/template-infrastructure/business/mid"
	"github.com/Celbux/template-infrastructure/business/user"
	"github.com/Celbux/template-infrastructure/foundation/web"
	"github.com/Celbux/template-infrastructure/thirdparty/datastore"
	"net/http"
	"os"
)

// API constructs a http.Handler with all application routes defined
func API(log i.Logger, dsClient *ds.Client, shutdown chan os.Signal, ) *web.App {

	app := web.NewApp(
		shutdown,
		mid.Logger(log),
		mid.Errors(log),
		mid.Namespace(log),
		mid.Metrics(),
		mid.Panics(log),
	)

	check := check{}

	// Create User Handlers
	u := User{
		Service: user.Service{
			Store: datastore.UserStore{
				DB: dsClient,
			},
			Log:  log,
		},
	}

	// Check Handlers
	app.Handle(http.MethodGet, "/readiness", check.readiness)
	app.Handle(http.MethodGet, "/liveness", check.liveness)

	// User Handlers
	app.Handle(http.MethodPost, "/createUser", u.createUser)
	app.Handle(http.MethodPost, "/getUser", u.getUser)

	return app

}
