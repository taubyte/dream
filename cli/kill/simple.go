package kill

import (
	"github.com/taubyte/dreamland/cli/command"
	"github.com/taubyte/dreamland/cli/common"
	client "github.com/taubyte/dreamland/service"
	"github.com/urfave/cli/v2"
)

func simple(multiverse *client.Client) *cli.Command {
	c := &cli.Command{
		Name:   "simple",
		Action: killSimple(multiverse),
	}

	// Attach gets
	command.NameWithDefault(c, common.DefaultClientName)
	command.Universe(c)

	return c
}

func killSimple(multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		universe := multiverse.Universe(c.String("universe"))
		return universe.KillSimple(c.String("name"))
	}
}
