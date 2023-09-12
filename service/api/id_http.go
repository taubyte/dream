package api

import (
	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream"
)

type UniverseInfo struct {
	Id string `json:"id"`
}

func (srv *multiverseService) idHttp() {
	srv.rest.GET(&httpIface.RouteDefinition{
		Path: "/id/{universe}",
		Vars: httpIface.Variables{
			Required: []string{"universe"},
		},
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			universeName, err := ctx.GetStringVariable("universe")
			if err != nil {
				return nil, err
			}

			u, err := libdream.GetUniverse(universeName)
			if err != nil {
				return nil, err
			}

			return UniverseInfo{Id: u.Id()}, nil
		},
	})
}
