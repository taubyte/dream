package inject

import (
	"fmt"

	commonIface "github.com/taubyte/go-interfaces/common"
)

func Service(name string, config *commonIface.ServiceConfig) Injectable {
	return Injectable{
		Name: name,
		Run: func(universe string) string {
			return fmt.Sprintf("/service/%s/%s", universe, name)
		},
		Config: config,
		Method: POST,
	}
}
