package configuration

type LazyRaven struct {
	Containers  []Container `yaml:"containers"`
	ParserCount int         `yaml:"parserCount"`
	Address     string      `yaml:"address"`
}

type Container struct {
	Id    string `yaml:"id"`
	Alias string `yaml:"alias"`
}
