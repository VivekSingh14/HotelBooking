package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfiguration
}

type ServerConfiguration struct {
	Host              string
	Port              string
	SSL               bool
	AppVersion        string
	CtxDefaultTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	Debug             bool
}

var c *Config
var configErr error

func GetConfig() (*Config, error) {
	c, configErr = loadConfig()
	return c, configErr
}

func loadConfig() (*Config, error) {
	os.Setenv("config", "local")
	test := os.Getenv("config")
	log.Println(test)
	configPath := getConfigPath(os.Getenv("config"))
	v := viper.New()
	v.SetConfigName(configPath)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	log.Println(v.ReadInConfig())

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	var c *Config
	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getConfigPath(env string) string {
	log.Printf("Loading config file for %s environment", env)
	if env == "local" {
		return "../config/config-env"
	}
	return "../config/config-env"
}
