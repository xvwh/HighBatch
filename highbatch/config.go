package highbatch

import (
	"github.com/BurntSushi/toml"
)

var Conf Config

type Config struct {
	Server ServerConfig
	Client ClientConfig
}

type ServerConfig struct {
	Name        string
	MongoDbHost string
}

type ClientConfigFile struct {
	Client ClientConfig
}

type ClientConfig struct {
	Tag    []string
	Master MasterConfig
}

type MasterConfig struct {
	Hostname string
	Port     string
}

func LoadConfig() (c Config) {
	Ld("in LoadConfig")
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		Le(err)
	}
	c = Conf
	return
}