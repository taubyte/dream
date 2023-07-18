package fixtures

import (
	"testing"

	_ "bitbucket.org/taubyte/auth/service"
	_ "bitbucket.org/taubyte/console/ui/service"
	moodyCommon "bitbucket.org/taubyte/go-moody-blues/common"
	_ "bitbucket.org/taubyte/monkey/service"
	_ "bitbucket.org/taubyte/patrick/service"
	_ "bitbucket.org/taubyte/tns/service"
	commonDreamland "github.com/taubyte/dreamland/core/common"
	dreamland "github.com/taubyte/dreamland/core/services"
	"github.com/taubyte/dreamland/helpers"
	commonIface "github.com/taubyte/go-interfaces/common"
	spec "github.com/taubyte/go-specs/common"
)

func TestImportProdProject(t *testing.T) {
	t.Skip("currently custom domains do not work on dreamland")
	moodyCommon.Dev = true

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
			"console": {},
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
