package provider

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ParamsProvider struct {
		ProviderURL string `yaml:"provider_url"`
		PrivateKey  string `yaml:"private_key"`
	} `yaml:"params_provider"`
}

var once sync.Once

func GetConfig() *Config {
	conf := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("config.yaml", conf); err != nil {
			log.Fatal(err)
		}
	})
	log.Println(conf.ParamsProvider)
	return conf
}
