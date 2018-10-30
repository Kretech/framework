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
			panic("load config file error" + err.Error())
		}

		err = conf.Load(b)
		if err != nil {
			panic(`parse toml error` + err.Error())
		}

		c.Singleton("config", conf)
		configx.NewConfigProxy(c)
	}

}
