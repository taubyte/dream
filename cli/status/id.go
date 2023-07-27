package status

import (
	"github.com/pterm/pterm"
	"github.com/taubyte/dreamland/cli/command"
	"github.com/taubyte/dreamland/cli/common"
	client "github.com/taubyte/dreamland/service"
	"github.com/taubyte/dreamland/service/api"
	"github.com/urfave/cli/v2"
)

func getID(multiverse *client.Client) *cli.Command {
	c := &cli.Command{
		Name:   "id",
		Action: getIDStatus(multiverse),
	}
	command.NameWithDefault(c, common.DefaultUniverseName)

	return c
}

func getIDStatus(multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		var id api.UniverseInfo
		id, err = multiverse.Universe(c.String("name")).Id()
		if err != nil {
			return
		}
		pterm.Success.Printf("Universe id: %s\n", id.Id)

		return
	}
}
