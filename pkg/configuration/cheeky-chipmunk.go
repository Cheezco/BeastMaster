package configuration

type CheekyChipmunk struct {
	LoggerPlugins []LoggerPlugin `yaml:"loggerPlugins"`
	PluginAddress string         `yaml:"pluginAddress"`
	Address       string         `yaml:"address"`
}

type LoggerPlugin struct {
	FileName      string `yaml:"fileName"`
	LaunchCommand string `yaml:"launchCommand"`
}
