package main

import (
	cheekychipmunk "BeastMaster/internal/cheeky-chipmunk"
	"BeastMaster/internal/debug"
	logging_plugins "BeastMaster/internal/logging-plugins"
	"BeastMaster/pkg"
	"BeastMaster/pkg/configuration"
	"flag"
	"log"
)

var configAddress string
var config configuration.Config

func main() {
	prepareFlags()

	err := cheekychipmunk.RegisterRpcServices()
	if err != nil {
		debug.Log(err)
		log.Fatal(err)
	}

	err = config.LoadConfig("tcp", configAddress)
	if err != nil {
		debug.Log(err)
		log.Fatal(err)
	}
	debug.Log("Config loaded")

	plugins := logging_plugins.DetectPlugins(config)
	debug.Logf("%d plugins detected. Failed to detect %d plugins", len(plugins), len(config.CheekyChipmunk.LoggerPlugins))
	logging_plugins.StartPlugins(plugins)

	go pkg.StartRpcServer("tcp", config.CheekyChipmunk.Address)
	debug.Logf("RPC server started. Listening at: %s", config.CheekyChipmunk.Address)

	select {}
}

func prepareFlags() {
	flag.StringVar(&configAddress, "ConfigAddress", "", "Configuration rpc server address")

	flag.Parse()
}
