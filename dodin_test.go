package dodin

import (
	"github.com/asarturas/dodin/cloud"
	"github.com/asarturas/dodin/cloud/fake"
	"testing"
)

func Test_dodin_returns_inventory_from_config_file_and_cloud_machines(t *testing.T) {
	configProvider := FakeConfigProvider{"[test]\nmembers=master[12]\n[test:vars]\nvariable=value"}
	cloudProvider := fake.Provider(
		cloud.GetCloudMachine("master3", "1.1.1.3"),
		cloud.GetCloudMachine("master1", "1.1.1.1"),
		cloud.GetCloudMachine("some-random-machine", "1.1.1.99"),
		cloud.GetCloudMachine("master2", "1.1.1.2"),
	)
	inventory := Dodin(configProvider, cloudProvider)
	expectedInventory := `{"test":{"hosts":["1.1.1.1","1.1.1.2"],"vars":{"variable":"value"}}}`
	if inventory.String() != expectedInventory {
		t.Errorf("Unexpected inventory (got, expected): \n %s \n %s", inventory.String(), expectedInventory)
	}
}
