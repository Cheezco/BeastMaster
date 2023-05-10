package logging_plugins

import (
	"BeastMaster/pkg/configuration"
	"errors"
	"fmt"
	"os"
)

func DetectPlugins(config configuration.Config) []configuration.LoggerPlugin {
	plugins := make([]configuration.LoggerPlugin, 0)
	for _, plugin := range config.CheekyChipmunk.LoggerPlugins {
		_, err := os.Stat("./LoggerPlugins/" + plugin.FileName + ".exe")
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		plugins = append(plugins, plugin)
	}

	return plugins
}

func StartPlugin(plugin configuration.LoggerPlugin) {
	fmt.Sprintln(plugin) // temporary hack
	//for {
	//	//scanner, cmd := internal.RunExecutableWithScanner(
	//	//	"./LoggerPlugins/"+plugin.FileName,
	//	//	"-Address="+plugin.Address,
	//	//)
	//	//
	//	//for scanner.Scan() {
	//	//	fmt.Println(scanner.Text())
	//	//}
	//	//err := cmd.Wait()
	//	//if err != nil {
	//	//	return
	//	//}
	//	//debug.Logf("%s logging plugin stopped. Restarting plugin...", plugin.FileName)
	//}

}

func StartPlugins(plugins []configuration.LoggerPlugin) {
	for _, plugin := range plugins {
		go StartPlugin(plugin)
	}
}
