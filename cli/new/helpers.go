package new

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/taubyte/dreamland/cli/common"
	client "github.com/taubyte/dreamland/service"
	"github.com/taubyte/dreamland/service/inject"
	commonIface "github.com/taubyte/go-interfaces/common"
	commonDreamland "github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
	slices "github.com/taubyte/utils/slices/string"
	"github.com/urfave/cli/v2"
)

func buildServiceConfig(enable, disable, binds []string) (map[string]commonIface.ServiceConfig, error) {
	_services := make([]string, 0)

	// Validation
	if len(disable) != 0 && len(enable) != 0 {
		return nil, errors.New("can't set enable and disable flags")
	}

	valid := services.ValidServices()
	if len(disable) > 0 {

		for _, s := range valid {
			disabled := false
			for _, d := range disable {
				if s == d {
					disabled = true
				}
			}
			if !disabled {
				_services = append(_services, s)
			}
		}

	} else if len(enable) > 0 {
		_services = append(_services, enable...)
	} else {
		_services = append(_services, valid...)
	}

	config, err := bindConfigServices(binds, _services)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func bindConfigServices(_binds []string, _services []string) (map[string]commonIface.ServiceConfig, error) {
	binds := make(map[string]map[string]int)
	for _, bind := range _binds {
		if len(bind) == 0 {
			continue
		}

		_def := strings.Split(bind, "@")
		if len(_def) != 1 && len(_def) != 2 {
			return nil, fmt.Errorf("processing bindings for `%s` failed", bind)
		}

		// grab the name
		name := _def[0]
		if len(name) == 0 || !slices.Contains(_services, name) {
			return nil, fmt.Errorf("could not bind port of service `%s`: disabled", name)
		}

		sub := ""

		// grab the port
		port := 0
		if len(_def) == 2 {
			_portDef := strings.Split(_def[1], "/")
			_port, _ := strconv.ParseInt(_portDef[0], 10, 32)
			port = int(_port)

			if len(_portDef) == 2 {
				valid_subs := common.ValidSubBinds
				sub = _portDef[1]
				if !slices.Contains(valid_subs, sub) {
					return nil, fmt.Errorf("`%s` not valid, should be one of: %s", sub, valid_subs)
				}
			} else {
				sub = "main"
			}
		}

		// define
		if _, exists := binds[name]; !exists {
			binds[name] = make(map[string]int)
		}
		if sub == "https" {
			binds[name]["secure"] = 1
			binds[name]["http"] = port
		} else {
			binds[name][sub] = port
		}

	}

	// Validate
	used := make(map[int]string, 0)
	for _service, portMap := range binds {
		for portIdx, port := range portMap {
			_used, ok := used[port]
			if ok {
				var errIdx string
				for idx, _port := range binds[_used] {
					if _port == port {
						errIdx = idx
					}
				}
				return nil, fmt.Errorf("attempted duplicate port bindings [%s@%d/%s] and [%s@%d/%s]", _used, port, errIdx, _service, port, portIdx)
			}
			used[port] = _service
		}
	}

	// gc
	used = nil

	config := make(map[string]commonIface.ServiceConfig, len(_services))
	for _, s := range _services {
		port := 0
		bind, ok := binds[s]
		if ok {
			port = bind["main"] // if not defined port == 0
		}

		config[s] = commonIface.ServiceConfig{
			CommonConfig: commonIface.CommonConfig{
				Disabled: false,
				Port:     port,
			},
			Others: bind,
		}
	}

	return config, nil
}

func buildConfig(c *cli.Context) (*commonDreamland.Config, error) {
	serviceConfig, err := buildServiceConfig(
		c.StringSlice("enable"),
		c.StringSlice("disable"),
		c.StringSlice("bind"),
	)
	if err != nil {
		return nil, err
	}

	simpleConfig, err := buildSimpleConfig(c.StringSlice("simples"))
	if err != nil {
		return nil, err
	}

	return &commonDreamland.Config{
		Services: serviceConfig,
		Simples:  simpleConfig,
	}, nil
}

func buildSimpleConfig(simples []string) (map[string]commonDreamland.SimpleConfig, error) {
	config := make(map[string]commonDreamland.SimpleConfig, len(simples))
	for _, simple := range simples {
		config[simple] = commonDreamland.SimpleConfig{
			Clients: getFilledClientConfig(),
		}
	}

	return config, nil
}

func runFixtures(c *cli.Context, multiverse *client.Client, universes []string) error {
	_fixtures := c.StringSlice("fixtures")
	fixtures := make([]inject.Injectable, 0)
	for _, fixture := range _fixtures {
		fixtures = append(fixtures, inject.Fixture(fixture, nil))
	}

	for _, universe := range universes {
		err := multiverse.Universe(universe).Inject(fixtures...)
		if err != nil {
			return fmt.Errorf("injecting fixtures into `%s` failed with: %w", universe, err)
		}
	}

	return nil
}

func startUniverses(c *cli.Context) (err error) {
	config, err := buildConfig(c)
	if err != nil {
		return err
	}

	for _, universe := range c.StringSlice("universes") {
		u := services.Multiverse(services.UniverseConfig{
			Name:     universe,
			Id:       c.String("id"),
			KeepRoot: c.Bool("keep"),
		})
		err = u.StartWithConfig(config)
		if err != nil {
			return err
		}
	}

	return
}

func startEmptyUniverses(c *cli.Context) (err error) {
	for _, universe := range c.StringSlice("universes") {
		u := services.Multiverse(services.UniverseConfig{Name: universe})
		err = u.StartWithConfig(&commonDreamland.Config{})
		if err != nil {
			return err
		}
	}

	return
}

func getFilledClientConfig() commonDreamland.SimpleConfigClients {
	return commonDreamland.SimpleConfigClients{
		Seer:      &commonIface.ClientConfig{},
		Auth:      &commonIface.ClientConfig{},
		Patrick:   &commonIface.ClientConfig{},
		TNS:       &commonIface.ClientConfig{},
		Monkey:    &commonIface.ClientConfig{},
		Hoarder:   &commonIface.ClientConfig{},
		Substrate: &commonIface.ClientConfig{},
	}
}
