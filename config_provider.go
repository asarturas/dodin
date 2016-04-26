package dodin

import "io/ioutil"

type ConfigProvider interface {
	Get() Config
}

type FakeConfigProvider struct {
	contents string
}

func (provider FakeConfigProvider) Get() Config {
	return ParseConfig(provider.contents)
}

type ConfigFileProvider struct {
	Filename string
}

func (provider ConfigFileProvider) Get() Config {
	contents, _ := ioutil.ReadFile(provider.Filename)
	return ParseConfig(string(contents[:]))
}
