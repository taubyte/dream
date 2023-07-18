package services

import (
	"fmt"

	"github.com/taubyte/dreamland/core/registry"
	commonIface "github.com/taubyte/go-interfaces/common"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	nodeIface "github.com/taubyte/go-interfaces/services/substrate"
)

func (u *Universe) CreateNodeService(config *commonIface.ServiceConfig) (peer.Node, error) {
	var err error

	if registry.Registry.Node.Service == nil {
		return nil, fmt.Errorf(`service is nil, have you imported _ "bitbucket.org/taubyte/node/service"`)
	}

	node, err := registry.Registry.Node.Service(u.ctx, config)
	if err != nil {
		return nil, err
	}

	_node, ok := node.(nodeIface.Service)
	if !ok {
		return nil, fmt.Errorf("failed type casting node into a service")
	}

	u.registerService("node", _node)
	u.toClose(_node)

	return _node.Node(), nil
}
