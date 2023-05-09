package container_monitoring

import (
	"BeastMaster/internal/monitoring"
	"BeastMaster/pkg"
	"BeastMaster/pkg/configuration"
)

type ContainerDataSender struct {
	ParsedDataChan          chan ContainerData
	ExporterSelectorAddress string
	CpuEnergyStorage        *monitoring.CpuEnergyStorage
}

func (s *ContainerDataSender) Run(config configuration.Config) error {
	for value := range s.ParsedDataChan {
		err := pkg.DialNoReply(
			"tcp",
			"localhost:1300",
			"ExportPluginService.Add", GetComputedData(value, s.CpuEnergyStorage.Value(), config.LazyRaven.Containers),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
