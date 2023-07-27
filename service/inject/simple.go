package inject

import (
	"fmt"

	commonDreamland "github.com/taubyte/tau/libdream/common"
)

func Simple(name string, config *commonDreamland.SimpleConfig) Injectable {
	return Injectable{
		Name: name,
		Run: func(universe string) string {
			return fmt.Sprintf("/simple/%s/%s", universe, name)
		},
		Config: config,
		Method: POST,
	}
}
