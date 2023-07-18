package common

import (
	"context"

	commonIface "github.com/taubyte/go-interfaces/common"
	"github.com/taubyte/go-interfaces/p2p/peer"
	authIface "github.com/taubyte/go-interfaces/services/auth"
	billingIface "github.com/taubyte/go-interfaces/services/billing"
	consoleIface "github.com/taubyte/go-interfaces/services/console"
	hoarderIface "github.com/taubyte/go-interfaces/services/hoarder"
	monkeyIface "github.com/taubyte/go-interfaces/services/monkey"
	patrickIface "github.com/taubyte/go-interfaces/services/patrick"
	qIface "github.com/taubyte/go-interfaces/services/q"
	seerIface "github.com/taubyte/go-interfaces/services/seer"
	nodeIface "github.com/taubyte/go-interfaces/services/substrate"
	tnsIface "github.com/taubyte/go-interfaces/services/tns"
)

type ClientCreationMethod func(*commonIface.ClientConfig) error

type SimpleConfig struct {
	commonIface.CommonConfig
	Clients SimpleConfigClients
}

type NodeInfo struct {
	Node  peer.Node
	Name  string
	Ports map[string]int
}

type SimpleConfigClients struct {
	Seer    *commonIface.ClientConfig
	Auth    *commonIface.ClientConfig
	Patrick *commonIface.ClientConfig
	TNS     *commonIface.ClientConfig
	Monkey  *commonIface.ClientConfig
	Hoarder *commonIface.ClientConfig
	Billing *commonIface.ClientConfig
	Node    *commonIface.ClientConfig
	Console *commonIface.ClientConfig
	Q       *commonIface.ClientConfig
}

type Config struct {
	Services map[string]commonIface.ServiceConfig
	Clients  map[string]commonIface.ClientConfig
	Simples  map[string]SimpleConfig
}
type Universe interface {
	Id() string
	Name() string
	Root() string // copy | just in case modified accidently
	Seer() seerIface.Service
	Auth() authIface.Service
	Patrick() patrickIface.Service
	TNS() tnsIface.Service
	Monkey() monkeyIface.Service
	Hoarder() hoarderIface.Service
	Billing() billingIface.Service
	Node() nodeIface.Service
	Console() consoleIface.Service
	Q() qIface.Service
	Context() context.Context
	Stop()
	// If no simple defined, starts one named StartAllDefaultSimple.
	StartAll(simples ...string) error
	Simple(name string) (Simple, error)
	StartWithConfig(mainConfig *Config) error
	Kill(serviceName string) error
	KillNodeByNameID(name string, id string) error
	GetPortHttp(peer.Node) (int, error)
	GetURLHttp(node peer.Node) (url string, err error)
	GetURLHttps(node peer.Node) (url string, err error)
	RunFixture(name string, params ...interface{}) error
	CreateSimpleNode(name string, config *SimpleConfig) (peer.Node, error)
	All() []peer.Node
	Register(node peer.Node, name string, ports map[string]int)
	Lookup(id string) (*NodeInfo, bool)
	Mesh(newNodes ...peer.Node)
	Service(name string, config *commonIface.ServiceConfig) error
	Provides(services ...string) error
	// Calls to grab services by pid
	SeerByPid(pid string) (seerIface.Service, bool)
	AuthByPid(pid string) (authIface.Service, bool)
	PatrickByPid(pid string) (patrickIface.Service, bool)
	TnsByPid(pid string) (tnsIface.Service, bool)
	MonkeyByPid(pid string) (monkeyIface.Service, bool)
	HoarderByPid(pid string) (hoarderIface.Service, bool)
	BillingByPid(pid string) (billingIface.Service, bool)
	NodeByPid(pid string) (nodeIface.Service, bool)
	ConsoleByPid(pid string) (consoleIface.Service, bool)
	QByPid(pid string) (qIface.Service, bool)
	ListNumber(name string) int
	GetServicePids(name string) ([]string, error)
}

type Simple interface {
	GetNode() peer.Node
	CreateSeerClient(config *commonIface.ClientConfig) error
	Seer() seerIface.Client
	CreateAuthClient(config *commonIface.ClientConfig) error
	Auth() authIface.Client
	CreatePatrickClient(config *commonIface.ClientConfig) error
	Patrick() patrickIface.Client
	CreateTNSClient(config *commonIface.ClientConfig) error
	TNS() tnsIface.Client
	CreateMonkeyClient(config *commonIface.ClientConfig) error
	Monkey() monkeyIface.Client
	CreateHoarderClient(config *commonIface.ClientConfig) error
	Hoarder() hoarderIface.Client
	CreateBillingClient(config *commonIface.ClientConfig) error
	Billing() billingIface.Client
	CreateConsoleClient(config *commonIface.ClientConfig) error
	Console() consoleIface.Client
	CreateQClient(config *commonIface.ClientConfig) error
	Q() qIface.Client
	Provides(clients ...string) error
}
