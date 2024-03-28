package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var config *Config

// readFile for reading development.yaml file
func readFile(cfg *Config) {
	f, err := os.Open("./config/development.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func BuildConfig() *Config {
	var cfg Config
	readFile(&cfg)
	config = &cfg
	// fmt.Printf("%+v", cfg)
	return &cfg
}

func GetConfig() *Config {
	return config
}

func (config *Config) DbURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.UserName,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)
}

func (config *Config) ServerURL() string {
	return fmt.Sprintf(
		"%s:%d",
		config.App.Host,
		config.App.Port,
	)
}
