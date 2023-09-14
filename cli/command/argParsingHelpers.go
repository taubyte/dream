package command

import "github.com/urfave/cli/v2"

func prependArgParsingToAction(c *cli.Command, argName string, argGetter func(ctx *cli.Context) (string, error)) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		arg, err := argGetter(ctx)
		if err != nil {
			return err
		}
		ctx.Set(argName, arg)

		// execute the original action at the end
		return c.Action(ctx)
		// ^I believe that this closure should have its own copy of c.Action before c.Action is changed, am I right?
	}
}

// detection of trailing unparsed flags might also go here
