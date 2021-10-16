package handlers

import (
	"github.com/Celbux/template-infrastructure/business/i"
	"github.com/Celbux/template-infrastructure/business/mid"
	"github.com/Celbux/template-infrastructure/foundation/web"
	"net/http"
	"os"
)

// API constructs an http.Handler with all application routes defined
func API(
	userHandlers UserHandlers,
	log i.Logger,
	shutdown chan os.Signal,
) *web.App {

	app := web.NewApp(
		shutdown,
		mid.Logger(log),
		mid.Errors(log),
		mid.Namespace(log),
		mid.Metrics(),
		mid.Panics(log),
	)

	check := check{}

	app.Handle(http.MethodGet, "/readiness", check.readiness)
	app.Handle(http.MethodGet, "/liveness", check.liveness)
	app.Handle(http.MethodPost, "/createUser", userHandlers.createUser)
	app.Handle(http.MethodPost, "/getUser", userHandlers.getUser)

	return app

}
