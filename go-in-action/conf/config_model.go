package conf

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
