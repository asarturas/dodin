package dodin

import (
	"encoding/json"
)

type InventoryGroup struct {
	Hosts []string          `json:"hosts"`
	Vars  map[string]string `json:"vars"`
}

type Inventory struct {
	groups map[string]InventoryGroup
}

func (inventory Inventory) String() string {
	result := ""
	numberOfGroups := len(inventory.groups)
	currentGroup := 1
	for name, group := range inventory.groups {
		groupJson, _ := json.Marshal(group)
		result += "\"" + name + "\":" + string(groupJson[:])
		if currentGroup != numberOfGroups {
			result += ","
			currentGroup++
		}
	}
	return "{" + result + "}"
}
