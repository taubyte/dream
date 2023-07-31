package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
)

func (srv *multiverseService) killUniverseHttp() {
	// Path to delete simples in a universe
	srv.rest.DELETE(&httpIface.RouteDefinition{
		Path: "/universe/{universe}",
		Vars: httpIface.Variables{
			Required: []string{"universe"},
		},
		Handler: srv.killUniverse,
	})
}

func (srv *multiverseService) killUniverse(ctx httpIface.Context) (interface{}, error) {
	// Grab the universe
	name, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, fmt.Errorf("failed getting name error %w", err)
	}

	if !srv.Exist(name) {
		return nil, fmt.Errorf("universe `%s` does not exist", name)
	}

	universe, err := srv.getUniverse(ctx)
	if err != nil {
		return nil, fmt.Errorf("killing universe `%s` failed with: %s", name, err.Error())
	}

	universe.Stop()

	return nil, nil
}
