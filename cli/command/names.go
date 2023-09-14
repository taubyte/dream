package command

import (
	"github.com/taubyte/dreamland/cli/flags"
	"github.com/urfave/cli/v2"
)

func Names(c *cli.Command) {
	attachNames(c, &flags.Names)
}

func attachNames(c *cli.Command, flag cli.Flag) {
	c.Flags = append(c.Flags, flag)

	if len(c.ArgsUsage) == 0 {
		c.ArgsUsage = "[name,...]"
	} else {
		c.ArgsUsage = "[name,...]" + c.ArgsUsage
	}

	c.Action = prependArgParsingToAction(c, "names", func(ctx *cli.Context) (string, error) {
		return getName(ctx)
	})

}
