package logging

import (
	"BeastMaster/pkg/configuration"
	"time"
)

type PluginService struct {
	LoggerPlugins []configuration.LoggerPlugin
}

func (s *PluginService) Log(log AppLog, _ *interface{}) error {
	log.Created = time.Now()
	//for _, plugin := range s.LoggerPlugins {
	//err := pkg.DialNoReply("tcp", plugin.Address, "LoggingPluginService.Log", log)
	//if err != nil {
	//	continue
	//}
	//}

	return nil
}
