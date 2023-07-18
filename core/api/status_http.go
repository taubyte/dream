package api

import (
	httpIface "github.com/taubyte/go-interfaces/services/http"
)

func statusHttp() {
	api.GET(&httpIface.RouteDefinition{
		Path:    "/status",
		Handler: statusHandler,
	})
}

func statusHandler(ctx httpIface.Context) (interface{}, error) {
	return mv.Status(), nil
}
