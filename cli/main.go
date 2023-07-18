package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	// Relative
	"github.com/taubyte/dreamland/cli/common"
	inject "github.com/taubyte/dreamland/cli/inject"
	"github.com/taubyte/dreamland/cli/kill"
	"github.com/taubyte/dreamland/cli/new"
	"github.com/taubyte/dreamland/cli/status"

	// Actual imports

	"bitbucket.org/taubyte/p2p/peer"
	"github.com/taubyte/dreamland/core/services"
	client "github.com/taubyte/dreamland/http"
	"github.com/urfave/cli/v2"

	// Empty imports for initializing fixtures, and client/service run methods
	_ "bitbucket.org/taubyte/auth/api/p2p"
	_ "bitbucket.org/taubyte/auth/service"
	_ "bitbucket.org/taubyte/billing/api/p2p"
	_ "bitbucket.org/taubyte/billing/service"
	_ "bitbucket.org/taubyte/console/api/p2p"
	_ "bitbucket.org/taubyte/console/ui/service"
	moodyCommon "bitbucket.org/taubyte/go-moody-blues/common"
	_ "bitbucket.org/taubyte/hoarder/api/p2p"
	_ "bitbucket.org/taubyte/hoarder/service"
	_ "bitbucket.org/taubyte/monkey/api/p2p"
	_ "bitbucket.org/taubyte/monkey/fixtures/compile"
	_ "bitbucket.org/taubyte/monkey/service"
	_ "bitbucket.org/taubyte/node/service"
	_ "bitbucket.org/taubyte/patrick/api/p2p"
	_ "bitbucket.org/taubyte/patrick/service"
	_ "bitbucket.org/taubyte/q-node/api/p2p"
	_ "bitbucket.org/taubyte/q-node/ui/service"
	_ "bitbucket.org/taubyte/seer-p2p-client"
	_ "bitbucket.org/taubyte/seer/service"
	_ "bitbucket.org/taubyte/tns-p2p-client"
	_ "bitbucket.org/taubyte/tns/service"
	_ "github.com/taubyte/dreamland/fixtures"
)

func main() {
	peer.DevMode = true
	moodyCommon.Dev = true
	ctx, ctxC := context.WithCancel(context.Background())

	defer func() {
		if common.DoDaemon {
			ctxC()
			services.Zeno()
		}
	}()

	multiverse, err := client.New(ctx, client.URL(common.DefaultDreamlandURL), client.Timeout(300*time.Second))
	if err != nil {
		log.Fatalf("Starting new dreamland client failed with: %s", err.Error())
	}

	err = defineCLI(&common.Context{Ctx: ctx, Multiverse: multiverse}).RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func defineCLI(ctx *common.Context) *(cli.App) {
	app := &cli.App{
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			new.Command(ctx),
			inject.Command(ctx),
			kill.Command(ctx),
			status.Command(ctx),
		},
		Suggest:              true,
		EnableBashCompletion: true,
	}

	return app
}

func setupTrap(ctxC context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for range c {
			ctxC()
			// sig is a ^C, handle it
		}
	}()
}
