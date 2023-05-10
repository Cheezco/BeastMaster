package configuration

type SleepyCapybara struct {
	Address       string         `yaml:"address"`
	PluginAddress string         `yaml:"exportAddress"`
	ExportPlugins []ExportPlugin `yaml:"exportPlugins"`
}

type ExportPlugin struct {
	FileName      string `yaml:"fileName"`
	LaunchCommand string `yaml:"launchCommand"`
}
