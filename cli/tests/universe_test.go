package tests

import (
	"fmt"
	"testing"
	"time"

	cliCommon "github.com/taubyte/dreamland/cli/common"

	common "github.com/taubyte/dreamland/core/common"
	dreamland "github.com/taubyte/dreamland/core/services"
	client "github.com/taubyte/dreamland/http"
	commonIface "github.com/taubyte/go-interfaces/common"
	_ "github.com/taubyte/odo/utils/dreamland"
)

var services = []string{"seer", "auth", "patrick", "tns", "monkey", "hoarder", "substrate"}

func TestKillService(t *testing.T) {
	t.Skip("this test needs to be redone")
	dreamland.BigBang()
	u := dreamland.Multiverse("KillService")
	err := u.StartWithConfig(&common.Config{
		Services: map[string]commonIface.ServiceConfig{},
		Clients:  map[string]commonIface.ClientConfig{},
		Simples:  map[string]common.SimpleConfig{},
	})

	if err != nil {
		t.Error(err)
		return
	}

	err = u.Service("tns", &commonIface.ServiceConfig{})
	if err != nil {
		t.Error(err)
		return
	}

	tnsIds, err := u.GetServicePids("tns")
	if err != nil {
		t.Error(err)
		return
	}
	idToDelete := tnsIds[0]

	err = u.KillNodeByNameID("tns", idToDelete)
	if err != nil {
		t.Error(err)
		return
	}

	tnsIds, err = u.GetServicePids("tns")
	if err != nil {
		t.Error(err)
		return
	}

	result := len(tnsIds)
	if result == 1 || result > 1 {
		t.Errorf("Service was not deleted with id: %s", idToDelete)
		return
	}

	multiverse, err := client.New(u.Context(), client.URL(cliCommon.DefaultDreamlandURL), client.Timeout(300*time.Second))
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := multiverse.Universe("KillService").Status()
	if err != nil {
		t.Error(err)
		return
	}

	if len(resp.Nodes) != 1 {
		t.Errorf("Service was not deleted with id: %s", idToDelete)
		return
	}
}

func TestKillSimple(t *testing.T) {
	testSimpleName := "client"
	universeName := "KillSimple"
	statusName := fmt.Sprintf("%s@%s", testSimpleName, universeName)

	dreamland.BigBang()
	u := dreamland.Multiverse(universeName)
	err := u.StartWithConfig(&common.Config{
		Clients: map[string]commonIface.ClientConfig{},
		Simples: map[string]common.SimpleConfig{
			testSimpleName: {},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	multiverse, err := client.New(u.Context(), client.URL(cliCommon.DefaultDreamlandURL), client.Timeout(1000*time.Second))
	if err != nil {
		t.Error(err)
		return
	}
	universeAPI := multiverse.Universe(universeName)

	simple, err := u.Simple(testSimpleName)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := universeAPI.Status()
	if err != nil {
		t.Error(err)
		return
	}
	var found bool
	for _, node := range resp.Nodes {
		if node.Name == statusName {
			found = true
		}
	}
	if found == false {
		t.Errorf("Couldn't find simple %s", testSimpleName)
		return
	}

	err = u.KillNodeByNameID("client", simple.PeerNode().ID().Pretty())
	if err != nil {
		t.Error(err)
		return
	}

	_, err = u.Simple("client")
	if err == nil {
		t.Error("Expected an error")
		return
	}

	resp, err = universeAPI.Status()
	if err != nil {
		t.Error(err)
		return
	}
	found = false
	for _, node := range resp.Nodes {
		if node.Name == statusName {
			found = true
		}
	}
	if found == true {
		t.Errorf("Found simple: %s when it should have been deleted", testSimpleName)
		return
	}

	// Create another with same name
	_, err = u.CreateSimpleNode("client", &common.SimpleConfig{
		CommonConfig: commonIface.CommonConfig{},
		Clients:      common.SimpleConfigClients{},
	})
	if err != nil {
		t.Error(err)
		return
	}

	resp, err = universeAPI.Status()
	if err != nil {
		t.Error(err)
		return
	}
	found = false
	for _, node := range resp.Nodes {
		if node.Name == statusName {
			found = true
		}
	}
	if found != true {
		t.Errorf("Couldn't find simple %s after recreating", testSimpleName)
		return
	}

}

func TestUniverseAll(t *testing.T) {
	u := dreamland.Multiverse("single")
	defer u.Stop()
	err := u.StartAll()
	if err != nil {
		t.Error(err)
		return
	}

	simple, err := u.Simple(common.StartAllDefaultSimple)
	if err != nil {
		t.Error(err)
		return
	}

	if u.Substrate() == nil {
		t.Error("Substrate is nil")
		return
	}
	if u.Seer() == nil || simple.Seer() == nil {
		t.Error("Seer || SeerClient == nil")
		return
	}
	if u.Auth() == nil || simple.Auth() == nil {
		t.Error("Auth || AuthClient == nil")
		return
	}
	if u.Patrick() == nil || simple.Patrick() == nil {
		t.Error("Patrick || PatrickClient == nil")
		return
	}
	if u.TNS() == nil || simple.TNS() == nil {
		t.Error("TNS || TNSClient == nil")
		return
	}
	if u.Monkey() == nil || simple.Monkey() == nil {
		t.Error("Monkey || MonkeyClient == nil")
		return
	}
	if u.Hoarder() == nil || simple.Hoarder() == nil {
		t.Error("Hoarder || HoarderClient == nil")
		return
	}

	// Wait for seer announce
	time.Sleep(time.Second * 5)
}

func TestMultipleServices(t *testing.T) {
	u := dreamland.Multiverse("multiple")
	defer u.Stop()
	err := u.StartWithConfig(&common.Config{
		Services: map[string]commonIface.ServiceConfig{
			"seer":    {Others: map[string]int{"copies": 1}},
			"auth":    {Others: map[string]int{"copies": 3}},
			"patrick": {Others: map[string]int{"copies": 3}},
			"tns":     {Others: map[string]int{"copies": 3}},
			"monkey":  {Others: map[string]int{"copies": 3}},
			"hoarder": {Others: map[string]int{"copies": 3}},
			"node":    {Others: map[string]int{"copies": 3}},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range services {
		if u.ListNumber(v) != 3 && v != "seer" {
			t.Errorf("Service %s does not have 2 copies got %d", v, u.ListNumber(v))
			return
		}
	}

	time.Sleep(time.Second * 1)
}
