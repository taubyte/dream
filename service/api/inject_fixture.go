package api

import (
	"errors"
	"fmt"

	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream/common"
)

func fixtureHttp() {
	serviceApi.POST(&httpIface.RouteDefinition{
		Path: "/fixture/{universe}/{fixture}",
		Vars: httpIface.Variables{
			Required: []string{"universe", "fixture", "params"},
		},
		Handler: apiHandlerFixture,
	})
}

func apiHandlerFixture(ctx httpIface.Context) (interface{}, error) {
	// Grab fixture to run
	fixture, err := ctx.GetStringVariable("fixture")
	if err != nil {
		return nil, fmt.Errorf("failed getting services error %w", err)
	}

	var found bool
	fixtures := mv.ValidFixtures()
	for _, _fixture := range fixtures {
		if fixture == _fixture {
			found = true
		}
	}
	if !found {
		return nil, fmt.Errorf("fixture `%s` not found in `%v`", fixture, fixtures)
	}

	// Grab the universe
	var universe common.Universe
	_name, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, fmt.Errorf("failed getting universe name error %w", err)
	}

	exist := mv.Exist(_name)
	if exist {
		universe = mv.Universe(_name)
	} else {
		return nil, fmt.Errorf("universe %s does not exist", _name)
	}

	params, ok := ctx.Variables()["params"].([]interface{})
	if !ok {
		return nil, errors.New("failed getting params")
	}

	err = universe.RunFixture(fixture, params...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
