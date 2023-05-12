package logging_plugins

import (
	"BeastMaster/internal"
	"BeastMaster/internal/debug"
	"BeastMaster/pkg/configuration"
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const basePluginPath = "./LoggingPlugins/"

func DetectPlugins(config configuration.Config) []configuration.LoggerPlugin {
	plugins := make([]configuration.LoggerPlugin, 0)
	for _, plugin := range config.CheekyChipmunk.LoggerPlugins {
		_, err := os.Stat(basePluginPath + plugin.FileName + ".exe")
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		plugins = append(plugins, plugin)
	}

	return plugins
}

func StartPlugin(plugin configuration.LoggerPlugin, serverAddress string) {
	for {
		var scanner *bufio.Scanner
		var cmd *exec.Cmd

		path := basePluginPath + plugin.FileName
		args := make([]string, 0)
		if plugin.LaunchCommand != "" {
			if strings.ToLower(plugin.LaunchCommand) == "python" {
				args = append(args,
					"-u",
					path,
					"-Address="+serverAddress,
				)

			} else {
				args = append(args,
					path,
					"-Address="+serverAddress,
				)
			}
			scanner, cmd = internal.RunExecutableWithScanner(plugin.LaunchCommand, args...)
		} else {
			scanner, cmd = internal.RunExecutableWithScanner(plugin.FileName, "-Address="+serverAddress)
		}

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		err := cmd.Wait()
		if err != nil {
			return
		}
		debug.Log("%s plugin stopped.")

	}

}

func StartPlugins(plugins []configuration.LoggerPlugin, serverAddress string) {
	for _, plugin := range plugins {
		go StartPlugin(plugin, serverAddress)
	}
}
