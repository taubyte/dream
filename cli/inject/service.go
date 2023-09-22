package inject

import (
	"github.com/taubyte/dreamland/cli/command"
	client "github.com/taubyte/dreamland/service"
	"github.com/taubyte/dreamland/service/inject"
	"github.com/taubyte/go-interfaces/common"
	specs "github.com/taubyte/go-specs/common"

	"github.com/urfave/cli/v2"
)

func service(multiverse *client.Client) []*cli.Command {
	validServices := specs.Protocols
	commands := make([]*cli.Command, len(validServices))

	for idx, _service := range validServices {
		c := &cli.Command{
			Name: _service,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name: "http",
				},
			},
			Action: runService(_service, multiverse),
		}
		command.Universe0(c)
		commands[idx] = c
	}

	return commands
}

func runService(name string, multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		universe := multiverse.Universe(c.String("universe"))

		others := make(map[string]int, 0)
		http := c.Int("http")
		if http != 0 {
			others["http"] = http
		}

		config := &common.ServiceConfig{
			Others: others,
		}

		return universe.Inject(inject.Service(name, config))
	}
}
