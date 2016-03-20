package dodin

import (
	"github.com/asarturas/dodin/cloud"
)

func Dodin(configProvider ConfigProvider, cloudProvider cloud.Provider) Inventory {
	config := configProvider.Get()
	machines := cloudProvider.GetMachines()
	inventory := BuildInventoryFrom(machines, config)
	return inventory
}
