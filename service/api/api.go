package api
//main landing page calling other services from here

//importing dependencies
import (
	"context"
	"time"

	goHttp "net/http"

	"github.com/pterm/pterm"
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
)

//creating a sstruct 
type multiverseService struct {
	rest httpIface.Service
	common.Multiverse
}

func BigBang() error {
	var err error

	//creating an instance
	srv := &multiverseService{
		Multiverse: services.NewMultiVerse(),
	}
	//listening to the server 
	srv.rest, err = http.New(srv.Context(), options.Listen(common.DreamlandApiListen), options.AllowedOrigins(true, []string{".*"}))
	if err != nil {
		return err
	}

	srv.setUpHttpRoutes().Start()

	waitCtx, waitCtxC := context.WithTimeout(srv.Context(), 10*time.Second)
	defer waitCtxC()
//infinite loop to report server status and check if the server is working or not 
	for {
		select {
		case <-waitCtx.Done():
			return waitCtx.Err()
		case <-time.After(100 * time.Millisecond):
			if srv.rest.Error() != nil {
				pterm.Error.Println("Dreamland failed to start")
				return srv.rest.Error()
			}
			_, err := goHttp.Get("http://" + common.DreamlandApiListen)
			if err == nil {
				pterm.Info.Println("Dreamland ready")
				return nil
			}
		}
	}
}
