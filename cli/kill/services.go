package kill

import (
	"strings"

	"github.com/taubyte/dreamland/cli/command"
	client "github.com/taubyte/dreamland/http"
	"github.com/urfave/cli/v2"
)

func services(multiverse *client.Client) *cli.Command {
	c := &cli.Command{
		Name:   "services",
		Action: killServices(multiverse),
	}

	command.Names(c)
	command.Universe(c)

	return c
}

func killServices(multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		universe := multiverse.Universe(c.String("universe"))
		services := strings.Split(c.String("names"), ",")

		for _, service := range services {
			err = universe.KillService(service)
			if err != nil {
				return err
			}
		}

		return
	}
}
