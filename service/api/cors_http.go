package api

import (
	"net/http"

	"github.com/taubyte/dreamland/core/api/cors"
	httpIface "github.com/taubyte/http"
)

func corsHttp() {
	api.LowLevel(&httpIface.LowLevelDefinition{
		Path: "/cors",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			cors.ProxyHandler(w, r)
		},
	})
}
