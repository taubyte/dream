package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
)

func killServiceHttp() {
	// Path to delete services/simple in a universe
	api.DELETE(&httpIface.RouteDefinition{
		Path: "/service/{universe}/{name}",
		Vars: httpIface.Variables{
			Required: []string{"universe", "name"},
		},
		Handler: killService,
	})
}

func killService(ctx httpIface.Context) (interface{}, error) {
	// FIXME: not stopping http server

	// Grab the universe
	universe, err := getUniverse(ctx)
	if err != nil {
		return nil, fmt.Errorf("killing service failed with: %s", err.Error())
	}

	// Grab services to kill
	_name, err := ctx.GetStringVariable("name")
	if err != nil {
		return nil, fmt.Errorf("failed getting services error %w", err)
	}

	// Kill services
	err = universe.Kill(_name)
	if err != nil {
		return nil, fmt.Errorf("failed killing %s with error: %w", _name, err)
	}

	return nil, nil
}
