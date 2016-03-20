package dodin

import "io/ioutil"

type ConfigProvider interface {
	Get() Config
}

type FakeConfigProvider struct {
	contents string
}

func (this FakeConfigProvider) Get() Config {
	return ParseConfig(this.contents)
}

type ConfigFileProvider struct {
	Filename string
}

func (this ConfigFileProvider) Get() Config {
	contents, _ := ioutil.ReadFile(this.Filename)
	return ParseConfig(string(contents[:]))
}
