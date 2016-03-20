package dodin

import (
	"testing"
)

func Test_intentory_string_representation_is_ansible_inventory_config(t *testing.T) {
	inventory := Inventory{
		groups: map[string]InventoryGroup{
			"test1": InventoryGroup{
				Hosts: []string{"host11", "host12"},
			},
			"test2": InventoryGroup{
				Hosts: []string{"host21", "host22"},
			},
		},
	}
	inventoryString := inventory.String()
	expectedString1 := `{"test1":{"hosts":["host11","host12"]},"test2":{"hosts":["host21","host22"]}}`
	expectedString2 := `{"test2":{"hosts":["host21","host22"]},"test1":{"hosts":["host11","host12"]}}`
	if inventoryString != expectedString1 && inventoryString != expectedString2 {
		t.Errorf("Unexpected inventory (got, expected): \n %s \n %s", inventoryString, expectedString1)
	}
}
