package fake

import "github.com/asarturas/dodin/cloud"

func Provider(machines... cloud.Machine) FakeProvider {
	return FakeProvider{
		machines: machines,
	}
}

type FakeProvider struct {
	machines []cloud.Machine
}

func (this FakeProvider) GetMachines() []cloud.Machine {
	return this.machines
}
