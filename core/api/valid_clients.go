package api

import (
	httpIface "github.com/taubyte/go-interfaces/services/http"
)

func validClients() {
	api.GET(&httpIface.RouteDefinition{
		Path:    "/spec/clients",
		Handler: clientsHandler,
	})
}

func clientsHandler(ctx httpIface.Context) (interface{}, error) {
	return mv.ValidClients(), nil
}
