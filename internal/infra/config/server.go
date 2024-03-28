package config

import "fmt"

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func ServerURL(config Config) string {
	return fmt.Sprintf(
		"%s:%d",
		config.App.Host,
		config.App.Port,
	)
}
