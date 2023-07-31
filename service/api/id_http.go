package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
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

			exists := srv.Exist(universeName)
			if !exists {
				return nil, fmt.Errorf("universe `%s` does not exit", universeName)
			}

			u := srv.Universe(universeName)
			return UniverseInfo{Id: u.Id()}, nil
		},
	})
}
