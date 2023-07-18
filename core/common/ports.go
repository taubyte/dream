package common

var (
	DefaultAuthPort    = 121
	DefaultHoarderPort = 142
	DefaultMonkeyPort  = 163
	DefaultPatrickPort = 184
	DefaultSeerPort    = 205
	DefaultTNSPort     = 226
	DefaultBillingPort = 248
	DefaultConsolePort = 260
	DefaultNodePort    = 282
	DefaultDnsPort     = 304
	DefaultQPort       = 326

	DefaultSeerHttpPort    = 403
	DefaultPatrickHttpPort = 424
	DefaultAuthHttpPort    = 445
	DefaultTNSHttpPort     = 466
	DefaultBillingHttpPort = 487
	DefaultConsoleHttpPort = 508
	DefaultNodeHttpPort    = 529
	DefaultQHttpPort       = 560

	DreamlandApiListen = DefaultURL + ":1421"
)

var (
	DefaultURL              string = "127.0.0.1"
	DefaultP2PListenFormat  string = "/ip4/" + DefaultURL + "/tcp/%d"
	DefaultHTTPListenFormat string = "%s://" + DefaultURL + ":%d"
)
