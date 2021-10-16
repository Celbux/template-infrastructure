module github.com/Celbux/template-infrastructure

go 1.16

// replace github.com/Celbux/dataapi => ../dataapi

require (
	cloud.google.com/go/bigquery v1.8.0
	cloud.google.com/go/cloudtasks v1.0.0
	cloud.google.com/go/datastore v1.6.0
	github.com/Celbux/dataapi v1.0.0
	github.com/ardanlabs/conf v1.5.0
	github.com/dimfeld/httptreemux v5.0.1+incompatible
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/google/uuid v1.3.0
	github.com/microcosm-cc/bluemonday v1.0.15
	github.com/pkg/errors v0.9.1
	google.golang.org/api v0.57.0
	gopkg.in/go-playground/validator.v9 v9.31.0
)
