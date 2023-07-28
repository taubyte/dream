package api

import (
	httpIface "github.com/taubyte/http"
)

func (srv *multiverseService) validFixtures() {
	srv.rest.GET(&httpIface.RouteDefinition{
		Path: "/spec/fixtures",
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			return srv.ValidFixtures(), nil
		},
	})
}
