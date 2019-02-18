package main

import (
	"log"

	"io/ioutil"
	"os"

	"github.com/juju/errors"
	"gopkg.in/yaml.v2"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Configuration struct {
	Importers []struct {
		Source      string `yaml: "source"`
		Destination string `yaml: "destination"`
	} `yaml: "importers"`
}

// Config read and memoize the config read from the file defined as
// MIAM_CONFIG_FILE of teh default is ./miam.yaml
func Config() *Configuration {
	var config_file string
	if os.Getenv("MIAM_CONFIG_FILE") != "" {
		config_file = os.Getenv("MIAM_CONFIG_FILE")
	} else {
		config_file = "miam.yaml"
	}
	data, err := ioutil.ReadFile(config_file)
	if err != nil {
		errors.Annotate(err, "error reading file"+config_file)
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
