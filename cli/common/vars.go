package common

import (
	dreamlandCommon "github.com/taubyte/dreamland/core/common"
)

var (
	DefaultDreamlandURL = "http://" + dreamlandCommon.DefaultURL + ":1421"
	BigBangStarted      = false
	DefaultUniverseName = "blackhole"
	DefaultClientName   = "client"
	DoDaemon            = false
	ValidSubBinds       = []string{"http", "p2p", "dns", "https", "verbose", "copies"}
)
