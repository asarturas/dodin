package dodin

import "github.com/asarturas/dodin/cloud"

func BuildInventoryFrom(machines []cloud.Machine, config Config) Inventory {
	inventory := Inventory{
		groups: make(map[string]InventoryGroup),
	}
	for _, groupConfig := range config.Groups() {
		group := InventoryGroup{}
		for _, machine := range machines {
			if groupConfig.MatchingMemberName(machine.Name()) {
				group.Hosts = append(group.Hosts, machine.IP())
			}
		}
		if len(group.Hosts) > 0 {
			inventory.groups[groupConfig.Name()] = group
		}
	}
	return inventory
}
