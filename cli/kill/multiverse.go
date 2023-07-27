package kill

import (
	client "github.com/taubyte/dreamland/service"
	"github.com/urfave/cli/v2"
)

func multiverse(multiverse *client.Client) *cli.Command {
	return &cli.Command{
		Name:   "multiverse",
		Action: killMultiverse(multiverse),
	}
}

func killMultiverse(multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		// TODO: this will simple kill a daemon if it exists

		return
	}
}
