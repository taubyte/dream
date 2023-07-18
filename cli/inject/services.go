package inject

import (
	"strings"

	"github.com/taubyte/dreamland/cli/command"
	client "github.com/taubyte/dreamland/http"
	"github.com/taubyte/dreamland/http/inject"
	"github.com/taubyte/go-interfaces/common"
	"github.com/urfave/cli/v2"
)

func services(multiverse *client.Client) *cli.Command {
	c := &cli.Command{
		Name:   "services",
		Action: runServices(multiverse),
	}

	command.Names(c)
	command.Universe(c)

	return c
}

func runServices(multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		universe := multiverse.Universe(c.String("universe"))
		config := &common.ServiceConfig{}

		services := strings.Split(c.String("names"), ",")

		injections := make([]inject.Injectable, 0)
		for _, service := range services {
			injections = append(injections, inject.Service(service, config))
		}

		return universe.Inject(injections...)
	}
}
