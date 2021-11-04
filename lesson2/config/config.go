package config

import (
	"encoding/json"
	"flag"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Name string `yaml:"name" json:"name"`
	Mode string `yaml:"mode" json:"mode"`
}

var (
	FlagName = flag.String("name", "", "names of people to greet")
	FlagMode = flag.String("mode", "", "hello or goodbye")
)

func (c *Configuration) ParseNames() []string {
	splitted := strings.Split(c.Name, ",")
	for i, s := range splitted {
		splitted[i] = strings.TrimSpace(s)
	}
	return splitted
}

func (c *Configuration) LoadConfig(file string) error {
	flag.Parse()
	err := godotenv.Load(file)
	if err != nil {
		return err
	}
	if *FlagName == "" {
		c.Name = os.Getenv("NAME")
	} else {
		c.Name = *FlagName
	}
	if *FlagMode == "" {
		c.Mode = os.Getenv("MODE")
	} else {
		c.Mode = *FlagMode
	}
	return nil
}

func (c *Configuration) ConfigFromJsonYaml(file string) error {
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if strings.Contains(file, ".json") {
		if err = json.Unmarshal(contents, &c); err != nil {
			return err
		}
	}
	if strings.Contains(file, ".yaml") {
		if err = yaml.Unmarshal(contents, &c); err != nil {
			return err
		}
	}
	return nil
}

func Load(file string) (Configuration, error) {
	var c Configuration
	if strings.Contains(file, ".json") || strings.Contains(file, ".yaml") {
		err := c.ConfigFromJsonYaml(file)
		if err != nil {
			return c, err
		}
	} else {
		err := c.LoadConfig(file)
		if err != nil {
			return c, err
		}
	}
	return c, nil
}
