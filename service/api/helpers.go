package api
//Importing dependencies
import (
	"fmt"

	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream/common"
)

//returning 
func (srv *multiverseService) getUniverse(ctx httpIface.Context) (common.Universe, error) {
	name, err := ctx.GetStringVariable("universe")
	//returning error if the fetching of universe failes
	if err != nil {
		return nil, fmt.Errorf("failed getting name with: %w", err)
	}
//if that name exists in the receiving service return an multiverse.Universe(name) with the paramter as name or return error
	exist := srv.Exist(name)
	if exist {
		return srv.Universe(name), nil
	} else {
		return nil, fmt.Errorf("universe %s does not exist", name)
	}
}
