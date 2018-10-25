package config

import (
	"io/ioutil"

	"github.com/Kretech/scaffold/config"
	"github.com/Kretech/scaffold/container"
)

type TomlServiceProvider struct {
	file string
}

func NewTomlServiceProvider() *TomlServiceProvider {
	return &TomlServiceProvider{
		file: "config.toml",
	}
}

func NewTomlServiceProviderWithFile(file string) *TomlServiceProvider {
	return &TomlServiceProvider{file: file}
}

func (p *TomlServiceProvider) Register(container *container.Container) {
	conf := new(config.Toml)

	b, err := ioutil.ReadFile(p.file)
	if err == nil {
		conf.Load(b)
	}

	container.BindInstance("config", conf)
}
