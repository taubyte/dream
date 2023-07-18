package services

import (
	"context"
	"sync"

	"github.com/taubyte/dreamland/core/common"
	commonIface "github.com/taubyte/go-interfaces/common"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	authIface "github.com/taubyte/go-interfaces/services/auth"
	consoleIface "github.com/taubyte/go-interfaces/services/console"
	hoarderIface "github.com/taubyte/go-interfaces/services/hoarder"
	monkeyIface "github.com/taubyte/go-interfaces/services/monkey"
	patrickIface "github.com/taubyte/go-interfaces/services/patrick"
	seerIface "github.com/taubyte/go-interfaces/services/seer"
	nodeIface "github.com/taubyte/go-interfaces/services/substrate"
	tnsIface "github.com/taubyte/go-interfaces/services/tns"
)

type Universe struct {
	ctx       context.Context
	ctxC      context.CancelFunc
	lock      sync.RWMutex
	name      string
	root      string
	id        string
	all       []peer.Node
	closables []commonIface.Service
	lookups   map[string]*common.NodeInfo
	portShift int
	service   map[string]*serviceInfo
	simples   map[string]*Simple

	keepRoot bool
}

type serviceInfo struct {
	nodes map[string]commonIface.Service
}

func (u *Universe) Name() string {
	return u.name
}

func (u *Universe) All() []peer.Node {
	return u.all
}

func (u *Universe) Lookup(id string) (*common.NodeInfo, bool) {
	u.lock.RLock()
	node, exist := u.lookups[id]
	u.lock.RUnlock()
	if !exist {
		return nil, false
	}
	return node, true
}

func (u *Universe) Root() string {
	return u.root
}

func (u *Universe) Context() context.Context {
	return u.ctx
}

func (u *Universe) Seer() seerIface.Service {
	ret, ok := first[seerIface.Service](u, u.service["seer"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) SeerByPid(pid string) (seerIface.Service, bool) {
	return byId[seerIface.Service](u, u.service["seer"].nodes, pid)
}

func (u *Universe) Auth() authIface.Service {
	ret, ok := first[authIface.Service](u, u.service["auth"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) AuthByPid(pid string) (authIface.Service, bool) {
	return byId[authIface.Service](u, u.service["auth"].nodes, pid)
}

func (u *Universe) Patrick() patrickIface.Service {
	ret, ok := first[patrickIface.Service](u, u.service["patrick"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) PatrickByPid(pid string) (patrickIface.Service, bool) {
	return byId[patrickIface.Service](u, u.service["patrick"].nodes, pid)
}

func (u *Universe) TNS() tnsIface.Service {
	ret, ok := first[tnsIface.Service](u, u.service["tns"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) TnsByPid(pid string) (tnsIface.Service, bool) {
	return byId[tnsIface.Service](u, u.service["tns"].nodes, pid)
}

func (u *Universe) Monkey() monkeyIface.Service {
	ret, ok := first[monkeyIface.Service](u, u.service["monkey"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) MonkeyByPid(pid string) (monkeyIface.Service, bool) {
	return byId[monkeyIface.Service](u, u.service["monkey"].nodes, pid)
}

func (u *Universe) Hoarder() hoarderIface.Service {
	ret, ok := first[hoarderIface.Service](u, u.service["hoarder"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) HoarderByPid(pid string) (hoarderIface.Service, bool) {
	return byId[hoarderIface.Service](u, u.service["hoarder"].nodes, pid)
}

func (u *Universe) Node() nodeIface.Service {
	ret, ok := first[nodeIface.Service](u, u.service["node"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) NodeByPid(pid string) (nodeIface.Service, bool) {
	return byId[nodeIface.Service](u, u.service["node"].nodes, pid)
}

func (u *Universe) Console() consoleIface.Service {
	ret, ok := first[consoleIface.Service](u, u.service["console"].nodes)
	if !ok {
		return nil
	}
	return ret
}

func (u *Universe) ConsoleByPid(pid string) (consoleIface.Service, bool) {
	return byId[consoleIface.Service](u, u.service["console"].nodes, pid)
}

func byId[T any](u *Universe, i map[string]commonIface.Service, pid string) (T, bool) {
	var result T
	u.lock.RLock()
	defer u.lock.RUnlock()
	a, ok := i[pid]
	if !ok {
		return result, false
	}
	_a, ok := a.(T)
	return _a, ok
}

func first[T any](u *Universe, i map[string]commonIface.Service) (T, bool) {
	var _nil T
	u.lock.RLock()
	defer u.lock.RUnlock()
	for _, s := range i {
		_s, ok := s.(T)
		if !ok || s == nil {
			return _nil, false
		}
		return _s, true
	}
	return _nil, false
}

func (u *Universe) ListNumber(name string) int {
	return len(u.service[name].nodes)
}
