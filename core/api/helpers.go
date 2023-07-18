package api

import (
	"fmt"

	"github.com/taubyte/dreamland/core/common"
	httpIface "github.com/taubyte/go-interfaces/services/http"
)

func getUniverse(ctx httpIface.Context) (common.Universe, error) {
	name, err := ctx.GetStringVariable("universe")
	if err != nil {
		return nil, fmt.Errorf("failed getting name with: %w", err)
	}

	exist := mv.Exist(name)
	if exist {
		return mv.Universe(name), nil
	} else {
		return nil, fmt.Errorf("universe %s does not exist", name)
	}
}
