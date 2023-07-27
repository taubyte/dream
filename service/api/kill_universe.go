package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
)

func killUniverseHttp() {
	// Path to delete simples in a universe
	serviceApi.DELETE(&httpIface.RouteDefinition{
		Path: "/universe/{universe}",
		Vars: httpIface.Variables{
			Required: []string{"universe"},
		},
		Handler: killUniverse,
	})
}

func killUniverse(ctx httpIface.Context) (interface{}, error) {
	// Grab the universe
	name, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, fmt.Errorf("failed getting name error %w", err)
	}

	if !mv.Exist(name) {
		return nil, fmt.Errorf("universe `%s` does not exist", name)
	}

	universe, err := getUniverse(ctx)
	if err != nil {
		return nil, fmt.Errorf("killing universe `%s` failed with: %s", name, err.Error())
	}

	universe.Stop()

	return nil, nil
}
