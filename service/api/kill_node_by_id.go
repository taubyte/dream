package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
)

func killNodeIdHttp() {
	// Path to delete services/simple in a universe
	serviceApi.DELETE(&httpIface.RouteDefinition{
		Path: "/node/{universe}/{name}/{id}",
		Vars: httpIface.Variables{
			Required: []string{"universe", "name", "id"},
		},
		Handler: killNodeById,
	})
}

func killNodeById(ctx httpIface.Context) (interface{}, error) {
	// Grab the universe
	universe, err := getUniverse(ctx)
	if err != nil {
		return nil, fmt.Errorf("killing service failed with: %s", err.Error())
	}

	// Grab node to kill
	_name, err := ctx.GetStringVariable("name")
	if err != nil {
		return nil, fmt.Errorf("failed getting name error %w", err)
	}

	_id, err := ctx.GetStringVariable("id")
	if err != nil {
		return nil, fmt.Errorf("failed getting id error %w", err)
	}

	// Kill node
	err = universe.KillNodeByNameID(_name, _id)
	if err != nil {
		return nil, fmt.Errorf("failed killing %s with error: %w", _id, err)
	}

	return nil, nil
}
