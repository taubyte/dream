package api

import (
	"github.com/taubyte/go-specs/common"
	httpIface "github.com/taubyte/http"
)

func (srv *multiverseService) validServices() {
	srv.rest.GET(&httpIface.RouteDefinition{
		Path: "/spec/services",
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			return common.Protocols, nil
		},
	})
}
