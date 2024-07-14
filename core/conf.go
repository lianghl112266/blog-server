package core

import (
	"blogServer/config"
	"blogServer/global"
	"os"

	"gopkg.in/yaml.v2"
)

func InitConf() error {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlConf, c)

	if err != nil {
		return err
	}

	global.Config = c
	return nil
}
