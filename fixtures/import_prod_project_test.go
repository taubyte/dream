package fixtures

import (
	"testing"

	commonDreamland "github.com/taubyte/dreamland/core/common"
	dreamland "github.com/taubyte/dreamland/core/services"
	"github.com/taubyte/dreamland/helpers"
	commonIface "github.com/taubyte/go-interfaces/common"
	spec "github.com/taubyte/go-specs/common"
	_ "github.com/taubyte/odo/protocols/auth"
	_ "github.com/taubyte/odo/protocols/monkey"
	_ "github.com/taubyte/odo/protocols/patrick"
	_ "github.com/taubyte/odo/protocols/tns"
)

func TestImportProdProject(t *testing.T) {
	t.Skip("currently custom domains do not work on dreamland")

	spec.DefaultBranch = "master_test"

	u := dreamland.Multiverse("TestImportProdProject")
	defer u.Stop()

	err := u.StartWithConfig(&commonDreamland.Config{
		Services: map[string]commonIface.ServiceConfig{
			"auth":    {},
			"tns":     {},
			"monkey":  {},
			"patrick": {},
			"hoarder": {},
		},
		Simples: map[string]commonDreamland.SimpleConfig{
			"client": {
				Clients: commonDreamland.SimpleConfigClients{
					TNS:     &commonIface.ClientConfig{},
					Auth:    &commonIface.ClientConfig{},
					Patrick: &commonIface.ClientConfig{},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = u.RunFixture("importProdProject", "QmYfMsCDvC9geoRRMCwRxvW1XSn3VQQoevBC48D9scmLJX", helpers.GitToken, "master_test")
	if err != nil {
		t.Error(err)
		return
	}
}
