package main

import (
	"log"

	"io/ioutil"
	"os"

	"github.com/juju/errors"
	"gopkg.in/yaml.v2"
)

// Configuration struct holding local miam config
type Configuration struct {
	Importers []struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	} `yaml:"importers"`
}

// Config read and memoize the config read from the file defined as
// MIAM_CONFIG_FILE of teh default is ./miam.yaml
func Config() *Configuration {
	var configFile string
	if os.Getenv("MIAM_CONFIG_FILE") != "" {
		configFile = os.Getenv("MIAM_CONFIG_FILE")
	} else {
		configFile = "miam.yaml"
	}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		errors.Annotate(err, "error reading file"+configFile)
		log.Fatalf("error: %v", err)
	}
	return parseConfig(string(data))
}

func parseConfig(data string) *Configuration {
	t := Configuration{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		errors.Annotate(err, "cannot parse yaml"+data)
		log.Fatalf("error: %v", err)
	}
	return &t
}
