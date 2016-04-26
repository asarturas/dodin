package dodin

import (
	"github.com/asarturas/dodin/cloud"
	"github.com/gogo/protobuf/test/group"
)

func FromMatchingMachines(config GroupConfig, machines []cloud.Machine) Group {
	newGroup := Group{
		name: config.Name(),
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
	name string
	members []cloud.Machine
	variables map[string]string
}

func (group Group) Name() string {
	return group.name
}

func (group Group) Members() []cloud.Machine {
	return group.members
}

func (group Group) Variables() map[string]string {
	return group.variables
}
