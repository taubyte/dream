package api

import (
	httpIface "github.com/taubyte/http"
)

func (srv *multiverseService) validClients() {
	srv.rest.GET(&httpIface.RouteDefinition{
		Path: "/spec/clients",
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			return srv.ValidClients(), nil
		},
	})
}
