package api

import (
	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream"
)

func (srv *multiverseService) validFixtures() {
	srv.rest.GET(&httpIface.RouteDefinition{
		Path: "/spec/fixtures",
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			return libdream.ValidFixtures(), nil
		},
	})
}
