package api

import (
	"time"

	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
)

var serviceApi httpIface.Service
var mv common.Multiverse

func BigBang() error {
	err := Start(services.NewMultiVerse())
	if err != nil {
		return err
	}
	return nil
}

func Start(m common.Multiverse) (err error) {
	mv = m
	serviceApi, err = http.New(m.Context(), options.Listen(common.DreamlandApiListen))
	if err != nil {
		return
	}

	setUpHttpRoutes()

	serviceApi.Start()

	time.Sleep(300 * time.Millisecond)

	err = serviceApi.Error()

	return
}
