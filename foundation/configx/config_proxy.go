package configx

import (
	"github.com/Kretech/contracts/config"
	"github.com/Kretech/contracts/container"
	"github.com/Kretech/framework/foundation"
)

type ConfigProxy struct {
	c container.Container
}

func (proxy *ConfigProxy) GetSub(key string) config.Repository {
	return proxy.c.Make("config").(config.Repository).GetSub(key)
}

func (proxy *ConfigProxy) Load(content []byte) error {
	return proxy.c.Make("config").(config.Repository).Load(content)
}

func (proxy *ConfigProxy) Get(key string) interface{} {
	return proxy.c.Make("config").(config.Repository).Get(key)
}

func (proxy *ConfigProxy) Config() config.Repository {
	return proxy.c.Make("config").(config.Repository)
}

// New ConfigProxy
func NewConfigProxy(c container.Container) *ConfigProxy {
	obj := &ConfigProxy{
		c: c,
	}
	return obj
}

var defaultConfigProxy = NewConfigProxy(foundation.App())

func Get(key string) interface{} {
	return defaultConfigProxy.Get(key)
}

func GetSub(key string) config.Repository {
	return defaultConfigProxy.GetSub(key)
}
