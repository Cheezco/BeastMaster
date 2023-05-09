package cheeky_chipmunk

import (
	"BeastMaster/pkg/logging"
	"net/rpc"
)

func RegisterRpcServices() error {
	loggerPluginService := new(logging.PluginService)

	err := rpc.Register(loggerPluginService)
	if err != nil {
		return err
	}

	return nil
}
