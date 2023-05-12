package container_monitoring

import (
	"BeastMaster/pkg/configuration"
	"encoding/json"
	"github.com/pbnjay/memory"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func ParseData(dataChan <-chan string, parsedDataChan chan<- ContainerData) error {
	for value := range dataChan {
		var parsedData ContainerData

		err := json.Unmarshal([]byte(value), &parsedData)
		if err != nil {
			continue
		}

		parsedDataChan <- parsedData
	}

	return nil
}

func GetComputedData(value ContainerData, cpuEnergyUsage int, containers []configuration.Container) ComputedContainerData {
	cpuDelta := value.CPUStats.CPUUsage.TotalUsage - value.PrecpuStats.CPUUsage.TotalUsage
	systemCpuDelta := value.CPUStats.SystemCPUUsage - value.PrecpuStats.SystemCPUUsage
	cpuUsage := (float64(cpuDelta) / float64(systemCpuDelta)) * 100.0

	usedMemory := value.MemoryStats.Usage - value.MemoryStats.Stats.Cache
	var memoryUsage float64
	if value.MemoryStats.Limit != 0 {
		memoryUsage = (float64(usedMemory) / float64(value.MemoryStats.Limit)) * 100.0
	} else {
		memoryUsage = 0
	}

	numberCpus := len(value.CPUStats.CPUUsage.PercpuUsage)
	cpuPercent, _ := cpu.Percent(0, false)
	cpuEnergy := (float64(cpuUsage) / 100) * float64(cpuEnergyUsage) * cpuPercent[0] / 100
	if cpuEnergy < 0 {
		cpuEnergy = 0
	}

	return ComputedContainerData{
		Name:               getContainerAlias(value.ID, value.Name[1:], containers),
		CpuDelta:           cpuDelta,
		SystemCpuDelta:     systemCpuDelta,
		CpuUsagePercent:    cpuUsage,
		UsedMemory:         usedMemory,
		MemoryUsagePercent: memoryUsage,
		MemoryEnergyUsed:   getRamEnergy(memoryUsage),
		NumberOfCpus:       numberCpus,
		CpuEnergyUsed:      cpuEnergy,
		TimeStamp:          time.Now(),
	}
}

// Super naive way of calculating ram power
func getRamEnergy(memoryUsagePercent float64) float64 {
	voltage := 1.35 // This only works for non overclocked ddr4, but I can't be asked to figure out how to get systems actual ram type
	totalMemory := memory.TotalMemory() / 1024 / 1024 / 1024
	maximumCurrentDraw := float64(totalMemory) * voltage / 8
	current := maximumCurrentDraw * memoryUsagePercent / 100

	return voltage * current
}

func getContainerAlias(id string, name string, containers []configuration.Container) string {
	for _, c := range containers {
		if c.Id != id || c.Alias == "" {
			continue
		}
		return c.Alias
	}

	return name
}
