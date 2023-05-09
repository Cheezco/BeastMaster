package configuration

type ConfigService struct {
	Config Config
}

func (c *ConfigService) Get(_ string, config *Config) error {
	config.BeastMaster = c.Config.BeastMaster
	config.LazyRaven = c.Config.LazyRaven
	config.CheekyChipmunk = c.Config.CheekyChipmunk
	config.SleepyCapybara = c.Config.SleepyCapybara

	return nil
}
