package services

import (
	"fmt"
	"time"

	"bitbucket.org/taubyte/p2p/keypair"
	peer "bitbucket.org/taubyte/p2p/peer"
	"github.com/taubyte/dreamland/core/common"
	commonIface "github.com/taubyte/go-interfaces/common"
	p2p "github.com/taubyte/go-interfaces/p2p/peer"
)

func ClientsWithDefaults(names ...string) common.SimpleConfigClients {
	config := common.SimpleConfigClients{}
	for _, name := range names {
		switch name {
		case "seer":
			config.Seer = &commonIface.ClientConfig{}
		case "auth":
			config.Auth = &commonIface.ClientConfig{}
		case "patrick":
			config.Patrick = &commonIface.ClientConfig{}
		case "tns":
			config.TNS = &commonIface.ClientConfig{}
		case "monkey":
			config.Monkey = &commonIface.ClientConfig{}
		case "hoarder":
			config.Hoarder = &commonIface.ClientConfig{}
		case "billing":
			config.Billing = &commonIface.ClientConfig{}
		case "console":
			config.Console = &commonIface.ClientConfig{}
		case "q":
			config.Q = &commonIface.ClientConfig{}
		}
	}
	return config
}

func (u *Universe) CreateSimpleNode(name string, config *common.SimpleConfig) (p2p.Node, error) {
	var err error

	if _, exists := u.simples[name]; exists {
		return nil, fmt.Errorf("simple Node `%s` exists in universe `%s`", name, u.Name())
	}

	if config.Port == 0 {
		config.Port = u.portShift + LastSimplePortAllocated()
	}

	simpleNode, err := peer.New(
		u.ctx,
		fmt.Sprintf("%s/simple-%s-%d", u.root, name, config.Port),
		keypair.NewRaw(),
		peer.DefaultSwarmKey(),
		[]string{fmt.Sprintf(common.DefaultP2PListenFormat, config.Port)},
		[]string{fmt.Sprintf(common.DefaultP2PListenFormat, config.Port)},
		false,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("failed creating me error: %v", err)
	}

	simple := &Simple{Node: simpleNode}

	all := simple.getAll()
	clientConfigs := map[string]*commonIface.ClientConfig{
		"seer":    config.Clients.Seer,
		"auth":    config.Clients.Auth,
		"patrick": config.Clients.Patrick,
		"tns":     config.Clients.TNS,
		"monkey":  config.Clients.Monkey,
		"hoarder": config.Clients.Hoarder,
		"billing": config.Clients.Billing,
		"console": config.Clients.Console,
		"q":       config.Clients.Q,
	}
	for clientType, config := range clientConfigs {
		if config == nil {
			continue
		}
		creationMethod, ok := all[clientType]
		if !ok {
			return nil, fmt.Errorf("unknown client type %s", clientType)
		}
		err = creationMethod(config)
		if err != nil {
			return nil, fmt.Errorf("client creation of %s failed with: %w", clientType, err)
		}
	}

	u.lock.Lock()
	u.simples[name] = simple
	u.lock.Unlock()

	time.Sleep(common.AfterStartDelay())
	u.Mesh(simpleNode)
	u.Register(simpleNode, name, map[string]int{"p2p": config.Port})

	return simpleNode, nil
}
