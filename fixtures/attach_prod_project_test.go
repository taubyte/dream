package fixtures

import (
	"testing"

	_ "bitbucket.org/taubyte/auth/service"
	moodyCommon "bitbucket.org/taubyte/go-moody-blues/common"
	_ "bitbucket.org/taubyte/hoarder/service"
	_ "bitbucket.org/taubyte/monkey/service"
	_ "bitbucket.org/taubyte/patrick/service"
	_ "bitbucket.org/taubyte/tns/service"
	commonDreamland "github.com/taubyte/dreamland/core/common"
	dreamland "github.com/taubyte/dreamland/core/services"
	"github.com/taubyte/dreamland/helpers"
	commonIface "github.com/taubyte/go-interfaces/common"
)

func TestAttachProdProject(t *testing.T) {
	moodyCommon.Dev = true

	u := dreamland.Multiverse("testrunlibrary")
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

	err = u.RunFixture("attachProdProject", helpers.ProjectID, helpers.GitToken, "dreamland")
	if err != nil {
		t.Error(err)
		return
	}

}
