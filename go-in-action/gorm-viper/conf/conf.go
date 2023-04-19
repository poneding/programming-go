package conf

import (
	"log"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type config struct {
	viper.Viper
}

func ReadFrom(file string) *config {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("read config failed: %s", err)
	}
	v := viper.GetViper()
	return &config{
		Viper: *v,
	}
}

func ReadFromOptions(opt ConfigOptions) *config {
	viper.AddConfigPath(opt.Path)
	viper.SetConfigType(opt.Type)
	viper.SetConfigName(opt.Name)

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("read config failed: %s", err)
	}
	v := viper.GetViper()
	return &config{
		Viper: *v,
	}
}

type ConfigOptions struct {
	Name string
	Type string
	Path string
}

func (c config) GetOrDefault(key string, def interface{}) interface{} {
	if !c.IsSet(key) {
		return def
	}
	return c.Get(key)
}

func (c config) GetUint8(key string) uint8 {
	return cast.ToUint8(c.Get(key))
}
