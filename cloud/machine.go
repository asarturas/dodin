package cloud

type Machine interface {
	Name() string
	IP() string
}

func GetCloudMachine(name, ip string) Machine {
	return CloudMachine{
		name: name,
		ip: ip,
	}
}

type CloudMachine struct {
	name, ip string
}

func (machine CloudMachine) Name() string {
	return machine.name
}

func (machine CloudMachine) IP() string {
	return machine.ip
}
