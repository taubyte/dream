package api

import (
	"fmt"

	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream/common"
)

func (srv *multiverseService) getUniverse(ctx httpIface.Context) (common.Universe, error) {
	name, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, fmt.Errorf("failed getting name with: %w", err)
	}

	exist := srv.Exist(name)
	if exist {
		return srv.Universe(name), nil
	} else {
		return nil, fmt.Errorf("universe %s does not exist", name)
	}
}
