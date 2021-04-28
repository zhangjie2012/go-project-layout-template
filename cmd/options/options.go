package options

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppOption struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var (
	appConfig *AppOption
)

// ParseOption parse config file and init `opt`
func ParseOption(fn string) (*AppOption, error) {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	opt := &AppOption{}
	if err := yaml.Unmarshal(bs, opt); err != nil {
		return nil, err
	}

	appConfig = opt

	return appConfig, err
}

// GetOption get server option
func GetOption() *AppOption {
	return appConfig
}
