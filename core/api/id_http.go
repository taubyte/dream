package api

import (
	"fmt"

	httpIface "github.com/taubyte/go-interfaces/services/http"
)

type UniverseInfo struct {
	Id string `json:"id"`
}

func idHttp() {
	api.GET(&httpIface.RouteDefinition{
		Path: "/id/{universe}",
		Vars: httpIface.Variables{
			Required: []string{"universe"},
		},
		Handler: apiHandlerId,
	})
}

func apiHandlerId(ctx httpIface.Context) (interface{}, error) {
	universeName, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, err
	}

	exists := mv.Exist(universeName)
	if !exists {
		return nil, fmt.Errorf("universe `%s` does not exit", universeName)
	}

	u := mv.Universe(universeName)
	return UniverseInfo{Id: u.Id()}, nil
}
