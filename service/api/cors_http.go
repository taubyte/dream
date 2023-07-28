package api

import (
	"net/http"

	"github.com/taubyte/dreamland/service/cors"
	httpIface "github.com/taubyte/http"
)

func (srv *multiverseService) corsHttp() {
	srv.rest.LowLevel(&httpIface.LowLevelDefinition{
		Path: "/cors",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			cors.ProxyHandler(w, r)
		},
	})
}
