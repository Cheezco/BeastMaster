package main

import (
	"BeastMaster/internal/debug"
	sleepycapybara "BeastMaster/internal/sleepy-capybara"
	"BeastMaster/pkg"
	"BeastMaster/pkg/configuration"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math"
)

var configAddress string
var config configuration.Config

func main() {
	prepareFlags()

	err := config.LoadConfig("tcp", configAddress)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	debug.Log("Config loaded")

	connections := make([]*websocket.Conn, 0)
	err = sleepycapybara.RegisterRpcServices(&connections)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	go func() {
		err := sleepycapybara.StartWebsocketListener(config.SleepyCapybara.PluginAddress, &connections)
		if err != nil {
			debug.Log(err)
			log.Fatal(err)
		}
	}()

	debug.Logf("Started websocket server. Listening at: %s", config.SleepyCapybara.PluginAddress)

	plugins := sleepycapybara.DetectPlugins(config)
	debug.Logf("%d plugins detected. Failed to detect %d plugins.", len(plugins),
		int(math.Max(float64(len(config.SleepyCapybara.ExportPlugins)-len(plugins)), 0)))
	sleepycapybara.StartPlugins(plugins, config.SleepyCapybara.PluginAddress)

	go pkg.StartRpcServer("tcp", config.SleepyCapybara.Address)
	debug.Logf("RPC server started. Listening at: %s", config.SleepyCapybara.Address)

	select {}
}

func prepareFlags() {
	flag.StringVar(&configAddress, "ConfigAddress", "", "Configuration rpc server address")

	flag.Parse()
}
