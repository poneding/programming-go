package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	SetConfigPath("./config.json")
	// conf.SetConfigPath("./config.yaml")
	c, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

type (
	ConfigModel struct {
		App AppModel `mapstructure:"app" json:"app"`
		DB  DBModel  `mapstructure:"db" json:"db"`
	}

	AppModel struct {
		About   string `mapstructure:"about" json:"about"`
		Version string `mapstructure:"version" json:"version"`
	}

	DBModel struct {
		Host       string `mapstructure:"host" json:"host"`
		Port       int32  `mapstructure:"port" json:"port"`
		Name       string `mapstructure:"name" json:"name"`
		User       string `mapstructure:"user" json:"user"`
		DBPassword string `mapstructure:"password" json:"password"`
	}
)

// type config struct {
// 	viper.Viper
// }

func loadConfigFrom(file string) (*ConfigModel, error) {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	v := viper.GetViper()
	// c := &config{Viper: *v}
	if err := v.Unmarshal(&configModel); err != nil {
		return nil, err
	}
	return configModel, nil
}

var (
	configPath  string
	configModel *ConfigModel
)

func LoadConfig() (*ConfigModel, error) {
	if configModel == nil {
		return loadConfigFrom(configPath)
	}
	return configModel, nil
}

func SetConfigPath(path string) {
	configPath = path
}
