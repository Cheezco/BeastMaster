package sleepy_capybara

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

func DetectPlugins(config configuration.Config) []configuration.ExportPlugin {
	plugins := make([]configuration.ExportPlugin, 0)
	for _, plugin := range config.SleepyCapybara.ExportPlugins {
		_, err := os.Stat("./ExportPlugins/" + plugin.FileName)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		plugins = append(plugins, plugin)
	}

	return plugins
}

func StartPlugin(plugin configuration.ExportPlugin, serverAddress string) {
	for {
		var scanner *bufio.Scanner
		var cmd *exec.Cmd

		path := "./ExportPlugins/" + plugin.FileName
		if plugin.LaunchCommand != "" {
			if strings.ToLower(plugin.LaunchCommand) == "python" {
				scanner, cmd = internal.RunExecutableWithScanner(
					plugin.LaunchCommand,
					"-u",
					path,
					"-Address="+serverAddress,
				)
			} else {
				scanner, cmd = internal.RunExecutableWithScanner(
					plugin.LaunchCommand,
					path,
					"-Address="+serverAddress,
				)
			}

		} else {
			scanner, cmd = internal.RunExecutableWithScanner(
				path,
				"-Address="+serverAddress,
			)
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

func startRequiredDockerCompose(path string) {
	internal.RunDockerComposeCommand("-f", path, "up", "-d")
}

func StartPlugins(plugins []configuration.ExportPlugin, serverAddress string) {
	for _, plugin := range plugins {
		if plugin.RequiredDockerCompose == "" {
			continue
		}
		startRequiredDockerCompose(plugin.RequiredDockerCompose)
	}
	for _, plugin := range plugins {
		go StartPlugin(plugin, serverAddress)
	}
}
