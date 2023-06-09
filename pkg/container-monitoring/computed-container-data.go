package container_monitoring

import "time"

type ComputedContainerData struct {
	Name                 string    `json:"name"`
	CpuDelta             int64     `json:"cpuDelta"`
	SystemCpuDelta       int64     `json:"systemCpuDelta"`
	CpuUsagePercent      float64   `json:"cpuUsagePercent"`
	UsedMemory           int64     `json:"usedMemory"`
	MemoryUsagePercent   float64   `json:"memoryUsagePercent"`
	MemoryEnergyUsed     float64   `json:"memoryEnergyUsed"`
	NumberOfCpus         int       `json:"numberOfCpus"`
	CpuEnergyUsed        float64   `json:"cpuEnergyUsed"`
	TimeStamp            time.Time `json:"timeStamp"`
	NetworkSentBytes     int       `json:"networkSentBytes"`
	NetworkReceivedBytes int       `json:"networkReceivedBytes"`
}
