package exporting

import (
	containerMonitoring "BeastMaster/pkg/container-monitoring"
	"github.com/gorilla/websocket"
)

type ExportPluginService struct {
	ExportConnections *[]*websocket.Conn
}

func (s *ExportPluginService) Add(data containerMonitoring.ComputedContainerData, _ *interface{}) error {
	for i := range *s.ExportConnections {
		conn := (*s.ExportConnections)[i]
		err := conn.WriteJSON(data)
		if err != nil {
			continue
		}
	}
	return nil
}
