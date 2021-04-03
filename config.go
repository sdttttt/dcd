package huck

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Counter []string
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func FromConfigFile(filename string) *Configuration {
	conf := NewConfiguration()

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln(err.Error())
	}

	yaml.Unmarshal(b, conf)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return conf
}
