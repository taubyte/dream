package api

import (
	"context"
	"time"

	goHttp "net/http"

	"github.com/pterm/pterm"
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream"
)

type multiverseService struct {
	rest httpIface.Service
	*libdream.Universe
}

func BigBang() error {
	var err error

	srv := &multiverseService{
		Universe: libdream.NewMultiVerse(),
	}

	srv.rest, err = http.New(srv.Context(), options.Listen(libdream.DreamlandApiListen), options.AllowedOrigins(true, []string{".*"}))
	if err != nil {
		return err
	}

	srv.setUpHttpRoutes().Start()

	waitCtx, waitCtxC := context.WithTimeout(srv.Context(), 10*time.Second)
	defer waitCtxC()

	for {
		select {
		case <-waitCtx.Done():
			return waitCtx.Err()
		case <-time.After(100 * time.Millisecond):
			if srv.rest.Error() != nil {
				pterm.Error.Println("Dreamland failed to start")
				return srv.rest.Error()
			}
			_, err := goHttp.Get("http://" + libdream.DreamlandApiListen)
			if err == nil {
				pterm.Info.Println("Dreamland ready")
				return nil
			}
		}
	}
}
