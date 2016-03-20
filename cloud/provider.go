package cloud

type Provider interface {
	GetMachines() []Machine
}
