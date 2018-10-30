package config

import (
	"io/ioutil"

	"github.com/Kretech/framework/foundation/configx"
	"github.com/Kretech/scaffold/config"
	"github.com/Kretech/scaffold/container"
)

func NewTomlFileServiceRegister(file string) func(*container.Container) {

	return func(c *container.Container) {
		conf := new(config.Toml)

		b, err := ioutil.ReadFile(file)
		if err != nil {
			panic("load config error" + err.Error())
		}
		conf.Load(b)

		c.Singleton("config", conf)
		configx.NewConfigProxy(c)
	}

}
