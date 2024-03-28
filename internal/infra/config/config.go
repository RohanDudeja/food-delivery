package config

type (
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"user_name"`
		DBName   string `yaml:"db_name"`
		Password string `yaml:"password"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}

	App struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Config struct {
		Database Database `yaml:"database"`
		Redis    Redis    `yaml:"redis"`
		App      App      `yaml:"app"`
	}
)
