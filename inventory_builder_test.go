package dodin

import (
	"testing"
	"github.com/asarturas/dodin/cloud"
)

func Test_inventory_builder_turns_dodin_config_and_cloud_machines_into_inventory(t *testing.T) {
	config := Config{
		groups: []GroupConfig{
			NewGroupConfig("test", "member[12]", map[string]string{"var1": "val1"}),
			NewGroupConfig("test2", "member9[23]", map[string]string{"variable2": "val2"}),
		},
	}
	machines := []cloud.Machine{
		cloud.GetCloudMachine("test1", "1.2.3.0"),
		cloud.GetCloudMachine("member1", "1.2.3.1"),
		cloud.GetCloudMachine("member9", "1.2.3.9"),
		cloud.GetCloudMachine("member92", "1.2.3.92"),
		cloud.GetCloudMachine("member2", "1.2.3.2"),
		cloud.GetCloudMachine("member93", "1.2.3.93"),
		cloud.GetCloudMachine("member94", "1.2.3.94"),
	}
	inventory := BuildInventoryFrom(machines, config)
	if len(inventory.groups) != 2 {
		t.Errorf("Expected to get %d groups, but got %d instead.", 2, len(inventory.groups))
	}
	if inventory.groups["test"].Hosts[0] != "1.2.3.1" || inventory.groups["test"].Hosts[1] != "1.2.3.2" {
		t.Error("Unexpected first group:", inventory)
	}
	if inventory.groups["test2"].Hosts[0] != "1.2.3.92" || inventory.groups["test2"].Hosts[1] != "1.2.3.93" {
		t.Error("Unexpected second group:", inventory)
	}
}
