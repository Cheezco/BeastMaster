package configuration

type CheekyChipmunk struct {
	LoggerPlugins []LoggerPlugin `yaml:"loggerPlugins"`
	Address       string         `yaml:"address"`
}

type LoggerPlugin struct {
	FileName string `yaml:"fileName"`
	Address  string `yaml:"address"`
}
