package configuration

type SleepyCapybara struct {
	Address       string         `yaml:"address"`
	ExportAddress string         `yaml:"exportAddress"`
	ExportPlugins []ExportPlugin `yaml:"exportPlugins"`
}

type ExportPlugin struct {
	FileName      string `yaml:"fileName"`
	LaunchCommand string `yaml:"launchCommand"`
}
