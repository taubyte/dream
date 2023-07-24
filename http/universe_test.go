package http

import (
	"context"
	"fmt"
	"testing"
	"time"

	commonDreamland "github.com/taubyte/dreamland/core/common"
	dreamland "github.com/taubyte/dreamland/core/services"
	"github.com/taubyte/dreamland/http/inject"
	commonIface "github.com/taubyte/go-interfaces/common"
	_ "github.com/taubyte/odo/protocols/auth"

	_ "github.com/taubyte/odo/protocols/hoarder"
	_ "github.com/taubyte/odo/protocols/monkey"
	_ "github.com/taubyte/odo/protocols/patrick"
	_ "github.com/taubyte/odo/protocols/seer"
	_ "github.com/taubyte/odo/protocols/tns"

	_ "github.com/taubyte/odo/clients/p2p/monkey"
	_ "github.com/taubyte/odo/clients/p2p/patrick"
	_ "github.com/taubyte/odo/clients/p2p/tns"
)

func TestRoutes(t *testing.T) {
	// start multiverse
	err := dreamland.BigBang()
	if err != nil {
		t.Errorf("Failed big bang with error: %v", err)
		return
	}

	u := dreamland.Multiverse("dreamland-http")
	defer u.Stop()

	err = u.StartWithConfig(&commonDreamland.Config{
		Services: map[string]commonIface.ServiceConfig{
			"monkey":  {},
			"auth":    {},
			"patrick": {},
			"seer":    {},
			"hoarder": {},
			"tns":     {},
		},
		Simples: map[string]commonDreamland.SimpleConfig{
			"client": {
				Clients: commonDreamland.SimpleConfigClients{
					Monkey:  &commonIface.ClientConfig{},
					Patrick: &commonIface.ClientConfig{},
					TNS:     &commonIface.ClientConfig{},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	ctx := context.Background()

	time.Sleep(2 * time.Second)

	client, err := New(ctx, URL("http://localhost:1421"), Timeout(60*time.Second))
	if err != nil {
		t.Errorf("Failed creating http client error: %v", err)
		return
	}

	universe := client.Universe("dreamland-http")

	// Create simple called test1
	err = universe.Inject(inject.Simple("test1", &commonDreamland.SimpleConfig{}))
	if err != nil {
		t.Errorf("Failed simples call with error: %v", err)
		return
	}

	time.Sleep(2 * time.Second)

	// Should not fail
	_, err = u.Simple("test1")
	if err != nil {
		t.Errorf("Failed getting simple with error: %v", err)
		return
	}

	// Should fail
	_, err = u.Simple("dne")
	if err == nil {
		t.Error("Should have failed, expecting to not find dne simple node")
		return
	}

	// Should not fail
	err = universe.KillService("seer")
	if err != nil {
		t.Errorf("Failed kill call with error: %v", err)
		return
	}

	// Should fail
	err = universe.KillService("seer")
	if err == nil {
		t.Error("Expected killing seer again to fail")
		return
	}

	// Should fail
	err = universe.Inject(inject.Fixture("should fail", "dne"))
	if err == nil {
		t.Error("Expecting fail for fixture not existing")
		return
	}

	test, err := client.Status()
	if err != nil {
		t.Error(err)
		return
	}
	_, ok := test["dreamland-http"]
	if ok == false {
		t.Error("Did not find universe in status")
		return
	}

	fmt.Println("-------------------END-------------")
}
