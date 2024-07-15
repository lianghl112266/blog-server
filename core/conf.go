package core

import (
	"blogServer/config"
	"blogServer/global"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "settings.yaml"

func InitConf() error {
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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
