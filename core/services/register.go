package services

import (
	"github.com/taubyte/dreamland/core/common"
	peer "github.com/taubyte/p2p/peer"
)

func (u *Universe) Register(node peer.Node, name string, ports map[string]int) {
	u.lock.Lock()
	defer u.lock.Unlock()
	u.lookups[node.ID().Pretty()] = &common.NodeInfo{
		Node:  node,
		Name:  name,
		Ports: ports,
	}
}
