package command

import (
	"errors"
	"strings"

	"github.com/taubyte/dreamland/cli/flags"
	"github.com/urfave/cli/v2"
)

// capitalized functions in go are public
// passing c which is a pointer to a value of the cli.Command type, and def which is a string
func NameWithDefault(c *cli.Command, def string) {

	// creates a new flags.Name which is a string flag from urfave/cli (cli.StringFlag)
	flag := flags.Name
	// according to urfave/cli's docs, DefaultText is the default text of the flag for usage purposes
	// def is being passed in the function to define the default text 
	flag.DefaultText = def
	// according to urfave/cli's cods, Value is default the value for this flag if not set by from any source
	// we're also using def for this
	flag.Value = def

	// calls attachName helper function which attaches the flag to the command
	attachName(c, &flag)
}

// capitalized functions in go are public
// passing c which is a pointer to a value of the cli.Command type
func Name(c *cli.Command) {
	// calls attachName helper function
	attachName(c, &flags.Name)
}

func attachName(c *cli.Command, flag cli.Flag) {
	// https://github.com/urfave/cli/blob/7b25ef5588871373af470ca0ef9d123691ecfeea/command.go#L27
	// the command struct has a property named Flags which holds a list of Flag structs
	// we're appending the flag to the list
	c.Flags = append(c.Flags, flag)

	// the command struct also holds ArgsUsage, a string which is a short description of the arguments
	if len(c.ArgsUsage) == 0 {
		// we're creating a new description if it doesn't exist
		c.ArgsUsage = "name"
	} else {
		// we're adding to the description if it does exist
		c.ArgsUsage = "name," + c.ArgsUsage
	}

	action := c.Action

	// modifies the action for the command (the function called when the command is invoked)
	c.Action = func(ctx *cli.Context) error {
		// gets the name from the getName function, verifying that it exists
		name, err := getName(ctx)
		if err != nil {
			return err
		}
		// modifies the context to hold this name
		ctx.Set("name", name)
		// keturns the original action with this context to keep the rest of the action's previous logic
		return action(ctx)
	}

}

// when name is args0 or flag -n this method will get
// or return an error
func getName(c *cli.Context) (name string, err error) {
	// I'm pretty new to golang and the concepts of Contexts
	// but this is pretty much just getting the first args of the context and ensuring it has a name assigned
	name = c.Args().First()
	if len(name) == 0 {
		name = c.String("name")
		if len(name) == 0 {
			err = errors.New("Please provide a name")
			return
		}
	} else {
		// this ensures the command is being written correctly
		if strings.HasPrefix(c.Args().Get(1), "-") {
			err = errors.New("Parse arguments failed: write [arguments] after -flags")
			return
		}
	}

	return
}
