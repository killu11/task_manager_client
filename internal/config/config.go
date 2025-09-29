package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Domain string `yaml:"domain"`
	Routes struct {
		User struct {
			SignUp Endpoint `yaml:"sign_up"`
			SignIn Endpoint `yaml:"sign_in"`
		} `yaml:"user"`

		Task struct {
			Create       Endpoint `yaml:"create"`
			GetByName    Endpoint `yaml:"get_task"`
			GetAll       Endpoint `yaml:"get_all"`
			UpdateStatus Endpoint `yaml:"update_status"`
			Delete       Endpoint `yaml:"delete"`
		} `yaml:"tasks"`
	} `yaml:"routes"`
}
type Endpoint struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

func NewConfig() *Config {
	f, err := os.Open("./internal/config/config.yaml")

	if err != nil {
		log.Fatalf("Не удалось загрузить конфиг: %v", err)
	}

	c := new(Config)
	if yaml.NewDecoder(f).Decode(c) != nil {
		log.Fatalln("Не удалось распарсить конфиг")
	}
	return c
}
