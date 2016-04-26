package dodin

import (
	"testing"
)

func Test_intentory_string_representation_is_ansible_inventory_config(t *testing.T) {
	inventory := Inventory{
		groups: map[string]InventoryGroup{
			"test1": {
				Hosts: []string{"host11", "host12"},
				Vars:  map[string]string{"variable": "value"},
			},
		},
	}
	inventoryString := inventory.String()
	expectedString := `{"test1":{"hosts":["host11","host12"],"vars":{"variable":"value"}}}`
	if inventoryString != expectedString {
		t.Errorf("Unexpected inventory (got, expected): \n %s \n %s", inventoryString, expectedString)
	}
}
