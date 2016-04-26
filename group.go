package dodin

import (
	"github.com/asarturas/dodin/cloud"
)

func FromMatchingMachines(config GroupConfig, machines []cloud.Machine) Group {
	newGroup := Group{
		name:      config.Name(),
		variables: config.Variables(),
	}
	for _, machine := range machines {
		if config.MatchingMemberName(machine.Name()) {
			newGroup.members = append(newGroup.members, machine)
		}
	}
	return newGroup
}

type Group struct {
	name      string
	members   []cloud.Machine
	variables map[string]string
}

func (this Group) Name() string {
	return this.name
}

func (this Group) Members() []cloud.Machine {
	return this.members
}

func (this Group) Variables() map[string]string {
	return this.variables
}
