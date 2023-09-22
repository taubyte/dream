package api

import (
	"github.com/taubyte/go-specs/common"
	httpIface "github.com/taubyte/http"
)

func (srv *multiverseService) validClients() {
	srv.rest.GET(&httpIface.RouteDefinition{
		Path: "/spec/clients",
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			return common.P2PStreamProtocols, nil
		},
	})
}
