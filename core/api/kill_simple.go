package api

import (
	"fmt"

	httpIface "github.com/taubyte/go-interfaces/services/http"
)

func killSimpleHttp() {
	// Path to delete simples in a universe
	api.DELETE(&httpIface.RouteDefinition{
		Path: "/simple/{universe}/{name}",
		Vars: httpIface.Variables{
			Required: []string{"universe", "name"},
		},
		Handler: killSimple,
	})
}

func killSimple(ctx httpIface.Context) (interface{}, error) {
	// Grab the universe
	universe, err := getUniverse(ctx)
	if err != nil {
		return nil, fmt.Errorf("killing simple failed with: %s", err.Error())
	}

	// Grab simple to kill
	_name, err := ctx.GetStringVariable("name")
	if err != nil {
		return nil, fmt.Errorf("failed getting simple to kill error %w", err)
	}

	// Kill simple
	err = universe.Kill(_name)
	if err != nil {
		return nil, fmt.Errorf("failed killing %s with error: %w", _name, err)
	}

	return nil, nil
}
