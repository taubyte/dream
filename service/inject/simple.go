package inject

import (
	"fmt"

	"github.com/taubyte/tau/libdream"
)

func Simple(name string, config *libdream.SimpleConfig) Injectable {
	return Injectable{
		Name: name,
		Run: func(universe string) string {
			return fmt.Sprintf("/simple/%s/%s", universe, name)
		},
		Config: config,
		Method: POST,
	}
}
