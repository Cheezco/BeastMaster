package main

import (
	"BeastMaster/internal/debug"
	"BeastMaster/internal/monitoring"
	"BeastMaster/pkg/configuration"
	container_monitoring "BeastMaster/pkg/container-monitoring"
	"flag"
	"fmt"
	"log"
)

var configAddress string
var config configuration.Config
var cpuEnergyStorage monitoring.CpuEnergyStorage

func main() {
	prepareFlags()

	err := config.LoadConfig("tcp", configAddress)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	debug.Log("Config loaded")

	cpuEnergyStorage = monitoring.CpuEnergyStorage{}
	go func() {
		err := monitoring.RunCpuMonitor(&cpuEnergyStorage)
		if err != nil {
			debug.Log(err)
			log.Fatal(err)
		}
	}()
	debug.Log("CPU monitor started")

	dataChan := make(chan string, len(config.LazyRaven.Containers)*2)
	parsedDataChan := make(chan container_monitoring.ContainerData, len(config.LazyRaven.Containers)*2)
	err = container_monitoring.MonitorContainers(config.LazyRaven, dataChan, parsedDataChan)
	if err != nil {
		debug.Log(err)
		log.Fatal(err)
	}
	debug.Log("Started monitoring containers")

	parsedDataSender := container_monitoring.ContainerDataSender{
		ParsedDataChan:          parsedDataChan,
		ExporterSelectorAddress: config.SleepyCapybara.Address,
		CpuEnergyStorage:        &cpuEnergyStorage,
	}
	go func() {
		err := parsedDataSender.Run(config)
		if err != nil {
			debug.Log(err)
			log.Fatal(err)
		}
	}()
	debug.Log("Data sender started")

	select {}
}

func prepareFlags() {
	flag.StringVar(&configAddress, "ConfigAddress", "", "Configuration rpc server address")

	flag.Parse()
}
