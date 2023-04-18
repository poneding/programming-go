package conf

import (
	"github.com/spf13/viper"
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
