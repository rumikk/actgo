package parser

import (
	"gopkg.in/yaml.v3"
)

type ConfigParser struct {
	ConfigFile string
}

type Process struct {
	Name      string    `yaml:"name"`
	Url       string    `yaml:"url"`
	Selector  string    `yaml:"selector"`
	Extractor string    `yaml:"extractor"`
	Actions   []*Action `yaml:"actions"`
}

type Action struct {
	Name string `yaml:"name"`
}

func NewProcessParser(processFile []byte) ([]*Process, error) {
	var processes []*Process
	err := yaml.Unmarshal(processFile, &processes)
	if err != nil {
		return nil, err
	}
	return processes, nil
}
