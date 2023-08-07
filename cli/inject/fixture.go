package inject

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/taubyte/dreamland/cli/command"
	client "github.com/taubyte/dreamland/service"
	"github.com/taubyte/dreamland/service/inject"
	dreamlandRegistry "github.com/taubyte/tau/libdream/registry"
	"github.com/urfave/cli/v2"
)

var noCamelRegEx = regexp.MustCompile(`(^|[a-z])([A-Z])`)

// turns camel-cased fixture name into something that looks better on command line
func noCamel(name string) string {
	ret := noCamelRegEx.ReplaceAllString(name, "${1}-${2}")
	ret = strings.ToLower(ret)
	return strings.TrimPrefix(ret, "-")
}

func fixture(multiverse *client.Client) []*cli.Command {
	commands := make([]*cli.Command, 0)

	var idx int
	for fixtureName, obj := range dreamlandRegistry.FixtureMap {
		if obj.BlockCLI {
			continue
		}

		c := &cli.Command{
			Name:        noCamel(fixtureName),
			Description: obj.Description,
			Usage:       obj.Description,
			Action:      runFixture(multiverse),
		}
		command.Universe0(c)

		for _, variable := range obj.Variables {
			aliases := []string{}
			if len(variable.Alias) != 0 {
				aliases = append(aliases, variable.Alias)
			}

			c.Flags = append(c.Flags, &cli.StringFlag{
				Name:     variable.Name,
				Usage:    variable.Description,
				Required: variable.Required,
				Aliases:  aliases,
			})
		}
		commands = append(commands, c)
		idx++
	}

	return commands
}

func runFixture(multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		universeName := c.String("universe")
		sendParams := make([]string, 0)
		for _, flag := range c.Command.Flags[1 : len(c.Command.Flags)-1] {
			sFlag := flag.(*cli.StringFlag)
			value := c.String(sFlag.Name)
			if len(value) == 0 && sFlag.Required {
				return fmt.Errorf("flag `%s` is required", sFlag.Name)
			}

			sendParams = append(sendParams, value)
		}

		universe := multiverse.Universe(universeName)
		return universe.Inject(inject.Fixture(c.Command.Name, sendParams))
	}
}
