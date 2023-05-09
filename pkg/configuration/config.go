package configuration

import (
	"BeastMaster/pkg"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	BeastMaster    BeastMaster    `yaml:"beastMaster"`
	LazyRaven      LazyRaven      `yaml:"lazyRaven"`
	CheekyChipmunk CheekyChipmunk `yaml:"cheekyChipmunk"`
	SleepyCapybara SleepyCapybara `yaml:"sleepyCapybara"`
}

func (c *Config) SetDefaultValues() {
	c.BeastMaster.ConfigAddress = "localhost:1000"

	c.LazyRaven.Address = "localhost:1100"
	c.LazyRaven.ParserCount = 1

	c.CheekyChipmunk.Address = "localhost:1200"

	c.SleepyCapybara.Address = "localhost:1300"
	c.SleepyCapybara.ExportAddress = "localhost:1400"
}

func (c *Config) LoadConfig(network, address string) error {
	err := pkg.DialNoArgs(network, address, "ConfigService.Get", c)
	return err
}

func (c *Config) LoadConfigLocal(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)

	return err
}
