package dodin

type Config struct {
	groups []GroupConfig
}

func (config Config) Groups() []GroupConfig {
	return config.groups
}
