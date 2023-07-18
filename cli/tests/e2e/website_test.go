package e2e

import (
	"fmt"
	"testing"
	"time"

	_ "bitbucket.org/taubyte/auth/service"
	_ "bitbucket.org/taubyte/hoarder/service"
	_ "bitbucket.org/taubyte/monkey/service"
	_ "bitbucket.org/taubyte/patrick/service"
	_ "bitbucket.org/taubyte/tns/service"
	commonDreamland "github.com/taubyte/dreamland/core/common"
	"github.com/taubyte/dreamland/core/services"
	commonTest "github.com/taubyte/dreamland/helpers"
	commonIface "github.com/taubyte/go-interfaces/common"
	specCommon "github.com/taubyte/go-specs/common"
	spec "github.com/taubyte/go-specs/methods"
)

var testWebsiteId = "2a547229-190d-412b-b13a-a4fb5306dec9"

func TestWebsite(t *testing.T) {
	u := services.Multiverse("single_e2e")
	defer u.Stop()

	err := u.StartWithConfig(&commonDreamland.Config{
		Services: map[string]commonIface.ServiceConfig{
			"monkey":  {},
			"hoarder": {},
			"tns":     {},
			"patrick": {},
			"auth":    {},
		},
		Simples: map[string]commonDreamland.SimpleConfig{
			"client": {
				Clients: commonDreamland.SimpleConfigClients{
					TNS:     &commonIface.ClientConfig{},
					Patrick: &commonIface.ClientConfig{},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = u.RunFixture("createProjectWithJobs")
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 15)

	mockAuthURL, err := u.GetURLHttp(u.Auth().Node())
	if err != nil {
		t.Error(err)
		return
	}

	mockPatrickURL, err := u.GetURLHttp(u.Patrick().Node())
	if err != nil {
		t.Error(err)
		return
	}

	err = commonTest.RegisterTestRepositories(u.Context(), mockAuthURL, commonTest.WebsiteRepo)
	if err != nil {
		t.Error(err)
		return
	}

	err = commonTest.PushJob(commonTest.WebsitePayload, mockPatrickURL, commonTest.WebsiteRepo)
	if err != nil {
		t.Error(err)
		return
	}

	// Check tns for website
	simple, err := u.Simple("client")
	if err != nil {
		t.Error(err)
		return
	}

	tnsClient := simple.TNS()

	attempts := 0
	var response interface{}

	fmt.Printf(`Getting from to
	projectID: %s
	webID: %s
	`, commonTest.ProjectID, testWebsiteId)

	commitObj, err := tnsClient.Fetch(specCommon.Current(commonTest.ProjectID, commonTest.Branch))
	if err != nil {
		t.Error(err)
		return
	}

	commit, ok := commitObj.Interface().(string)
	if ok == false {
		t.Errorf("Could not convert commit object interface{} `%v` to string", commitObj.Interface())
		return
	}

	assetPath, err := spec.GetTNSAssetPath(commonTest.ProjectID, testWebsiteId, commit)
	if err != nil {
		t.Error(err)
		return
	}

	for {
		response, err = tnsClient.Fetch(assetPath)
		if err != nil {
			t.Error(err)
			return
		}
		if response != nil {
			fmt.Println("Response from TNS", response)
			break
		}
		attempts += 1
		if attempts == 60 {
			t.Errorf("Failed fetching from tns after %d attempts", attempts)
			return
		}
		time.Sleep(1 * time.Second)
	}

	// buildCID, ok := response.(string)

}
