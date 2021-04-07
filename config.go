package huck

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var conf *Configuration

// Configuration is the object structure of huck.yml.
type Configuration struct {
	Counter []string
}

// NewConfiguration to create a empty Configuration struct.
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// FromConfigFile from huck.yml read data to Configuration struct.
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
