package common

import (
	"github.com/taubyte/tau/libdream"
)

var (
	DefaultDreamlandURL = "http://" + libdream.DreamlandApiListen
	DefaultUniverseName = "blackhole"
	DefaultClientName   = "client"
	DoDaemon            = false
	ValidSubBinds       = []string{"http", "p2p", "dns", "https", "verbose", "copies"}
)
