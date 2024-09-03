package conf

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	ElasticsearchAddress []string `yaml:"elasticsearch_address"`
	UserName             string   `yaml:"user_name"`
	Password             string   `yaml:"password"`
	IndexPrefix          string   `yaml:"index_prefix"`
	ServerAddress        string   `yaml:"server_address"`
	ServerPort           int64    `yaml:"server_port"`
}

func (c *Config) Load(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return errors.New("should specify a config file")
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}
