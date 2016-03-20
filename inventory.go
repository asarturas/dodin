package dodin

import (
	"encoding/json"
)

type InventoryGroup struct {
	Hosts []string `json:"hosts"`
	Vars map[string]string `json:"vars"`
}

type Inventory struct {
	groups map[string]InventoryGroup
}

func (this Inventory) String() string {
	result := ""
	numberOfGroups := len(this.groups)
	currentGroup := 1
	for name, group := range this.groups {
		groupJson, _ := json.Marshal(group)
		result += "\"" + name + "\":" + string(groupJson[:])
		if currentGroup != numberOfGroups {
			result += ","
			currentGroup++
		}
	}
	return "{" + result + "}"
}
