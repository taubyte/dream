package services

import (
	"fmt"

	"github.com/taubyte/dreamland/core/registry"
	commonIface "github.com/taubyte/go-interfaces/common"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	tnsIface "github.com/taubyte/go-interfaces/services/tns"
)

func (u *Universe) CreateTNSService(config *commonIface.ServiceConfig) (peer.Node, error) {
	if registry.Registry.TNS.Service == nil {
		return nil, fmt.Errorf(`Service is nil, have you imported _ "github.com/taubyte/odo/protocols/tns/service"`)
	}

	tns, err := registry.Registry.TNS.Service(u.ctx, config)
	if err != nil {
		return nil, err
	}

	_tns, ok := tns.(tnsIface.Service)
	if !ok {
		return nil, fmt.Errorf("failed type casting tns into a service")
	}

	u.registerService("tns", _tns)
	u.toClose(_tns)

	return _tns.Node(), nil
}

func (s *Simple) CreateTNSClient(config *commonIface.ClientConfig) error {
	if registry.Registry.TNS.Client == nil {
		return fmt.Errorf(`client is nil, have you imported _ "github.com/taubyte/odo/clients/p2p/tns"`)
	}
	_tns, err := registry.Registry.TNS.Client(s.Node, config)
	if err != nil {
		return err
	}

	var ok bool
	s.Clients.tns, ok = _tns.(tnsIface.Client)
	if !ok {
		return fmt.Errorf("setting tns client failed, not OK")
	}

	return nil
}
