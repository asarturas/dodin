package cloud

type Machine interface {
	Name() string
	IP() string
}

func GetCloudMachine(name, ip string) Machine {
	return CloudMachine{
		name: name,
		ip:   ip,
	}
}

type CloudMachine struct {
	name, ip string
}

func (this CloudMachine) Name() string {
	return this.name
}

func (this CloudMachine) IP() string {
	return this.ip
}
