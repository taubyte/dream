package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream/common"
)

func injectUniverseHttp() {
	// Path to create simples in a universe
	serviceApi.POST(&httpIface.RouteDefinition{
		Path: "/universe/{universe}",
		Vars: httpIface.Variables{
			Required: []string{"universe", "config"},
		},
		Handler: apiHandlerUniverse,
	})
}

func apiHandlerUniverse(ctx httpIface.Context) (interface{}, error) {
	name, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, fmt.Errorf("failed getting name with: %w", err)
	}

	// Grab the universe
	exist := mv.Exist(name)
	if exist {
		return nil, fmt.Errorf("universe `%s` already exists", name)
	}

	config := struct {
		Config *common.Config
	}{}

	err = ctx.ParseBody(&config)
	if err != nil {
		return nil, err
	}

	u := mv.Universe(name)
	return nil, u.StartWithConfig(config.Config)
}
