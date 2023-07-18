package services

import (
	"fmt"

	"github.com/taubyte/dreamland/core/registry"
	commonIface "github.com/taubyte/go-interfaces/common"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	qIface "github.com/taubyte/go-interfaces/services/q"
)

func (u *Universe) CreateQService(config *commonIface.ServiceConfig) (peer.Node, error) {
	var err error

	if registry.Registry.Q.Service == nil {
		return nil, fmt.Errorf(`service is nil, have you imported _ "bitbucket.org/taubyte/q-node/ui/service"`)
	}

	q, err := registry.Registry.Q.Service(u.ctx, config)
	if err != nil {
		return nil, fmt.Errorf("calling Q service registry method failed with: %s", err)
	}

	_q, ok := q.(qIface.Service)
	if !ok {
		return nil, fmt.Errorf("got q service type(%T) expected(%T)", q, *new(qIface.Service))
	}

	u.registerService("q", _q)
	u.toClose(_q)

	return _q.Node(), nil
}

func (s *Simple) CreateQClient(config *commonIface.ClientConfig) error {
	if registry.Registry.Q.Client == nil {
		return fmt.Errorf(`client is nil, have you imported _ "bitbucket.org/taubyte/q-node/api/p2p"`)
	}

	_q, err := registry.Registry.Q.Client(s.Node, config)
	if err != nil {
		return fmt.Errorf("calling Q client registry method failed with: %s", err)
	}

	var ok bool
	s.Clients.q, ok = _q.(qIface.Client)
	if !ok {
		return fmt.Errorf("got q client type(%T) expected(%T)", _q, *new(qIface.Client))
	}

	return nil

}
