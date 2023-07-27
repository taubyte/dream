package common

import (
	dreamlandCommon "github.com/taubyte/tau/libdream/common"
)

var (
	DefaultDreamlandURL = "http://" + dreamlandCommon.DreamlandApiListen
	BigBangStarted      = false
	DefaultUniverseName = "blackhole"
	DefaultClientName   = "client"
	DoDaemon            = false
	ValidSubBinds       = []string{"http", "p2p", "dns", "https", "verbose", "copies"}
)
