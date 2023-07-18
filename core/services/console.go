package services

import (
	"fmt"

	"github.com/taubyte/dreamland/core/common"
	"github.com/taubyte/dreamland/core/registry"
	commonIface "github.com/taubyte/go-interfaces/common"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	consoleIface "github.com/taubyte/go-interfaces/services/console"
)

func (u *Universe) CreateConsoleService(config *commonIface.ServiceConfig) (peer.Node, error) {
	var err error

	authPort := config.Others["auth"]
	if authPort == 0 {
		authPort := u.portShift + common.DefaultAuthHttpPort
		config.Others["auth"] = authPort
	}

	patrickPort := config.Others["patrick"]
	if patrickPort == 0 {
		patrickPort := u.portShift + common.DefaultAuthHttpPort
		config.Others["patrick"] = patrickPort
	}

	if registry.Registry.Console.Service == nil {
		return nil, fmt.Errorf(`Service is nil, have you imported _ "bitbucket.org/taubyte/console/ui/service"`)
	}

	console, err := registry.Registry.Console.Service(u.ctx, config)
	if err != nil {
		return nil, err
	}

	_console, ok := console.(consoleIface.Service)
	if !ok {
		return nil, fmt.Errorf("failed type casting console into a service")
	}

	u.registerService("console", _console)
	u.toClose(_console)

	return _console.Node(), nil
}

func (s *Simple) CreateConsoleClient(config *commonIface.ClientConfig) error {
	if registry.Registry.Console.Client == nil {
		return fmt.Errorf("client is nil, have you imported _ \"bitbucket.org/taubyte/console/api/p2p\"")
	}

	_console, err := registry.Registry.Console.Client(s.Node, config)
	if err != nil {
		return err
	}

	var ok bool
	s.Clients.console, ok = _console.(consoleIface.Client)
	if !ok {
		return fmt.Errorf("setting console client failed, not OK")
	}

	return nil

}
