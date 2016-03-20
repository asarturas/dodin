package dodin

type Config struct {
	groups []GroupConfig
}

func (this Config) Groups() []GroupConfig {
	return this.groups
}
