package main

import (
	"BeastMaster/internal/beast-master"
	"BeastMaster/internal/debug"
	"BeastMaster/pkg"
	"BeastMaster/pkg/configuration"
	"flag"
)

var configPath string
var config configuration.Config

func main() {
	prepareFlags()

	config.SetDefaultValues()
	err := config.LoadConfigLocal("defaultConfig.yml")
	if err != nil {
		debug.Log(err)
		panic(err)
	}
	debug.Log("Config loaded")

	err = beast_master.RegisterRpcServices(config)
	go pkg.StartRpcServer("tcp", config.BeastMaster.ConfigAddress)
	debug.Logf("RPC server started. Listening at: %s", config.BeastMaster.ConfigAddress)

	beast_master.StartModules(config)

	select {}
}

func prepareFlags() {
	flag.StringVar(&configPath, "Config", "defaultConfig.yml", "Configuration file path.")

	flag.Parse()
}
