package handlers

import (
	"net/http"
	"os"

	"github.com/Celbux/template-infrastructure/business/i"
	"github.com/Celbux/template-infrastructure/business/mid"
	"github.com/Celbux/template-infrastructure/foundation/web"
)

// API constructs an http.Handler with all application routes defined
func API(
	dataapi DataAPIHandlers,
	log i.Logger,
	shutdown chan os.Signal,
) *web.App {

	app := web.NewApp(
		shutdown,
		mid.Logger(log),
		mid.Errors(log),
		mid.Metrics(),
		mid.Panics(log),
	)

	check := check{}

	app.Handle(http.MethodGet, "/readiness", check.readiness)
	app.Handle(http.MethodGet, "/liveness", check.liveness)
	app.Handle(http.MethodPost, "/evaluate/ns/:ns", dataapi.evaluateHandler)
	app.Handle(http.MethodPost, "/evaluate", dataapi.evaluateHandler)

	return app

}
