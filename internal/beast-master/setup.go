package beast_master

import (
	"BeastMaster/internal"
	"BeastMaster/internal/debug"
	"BeastMaster/pkg/configuration"
	"fmt"
	"net/rpc"
)

func StartModules(config configuration.Config) {
	cheekyChipmunkStarted := false
	sleepyCapybaraStarted := false
	lazyRavenStarted := false
	go startModule(
		"./CheekyChipmunk.exe",
		&cheekyChipmunkStarted,
		"-ConfigAddress="+config.BeastMaster.ConfigAddress,
	)
	for !cheekyChipmunkStarted {
	}

	go startModule(
		"./SleepyCapybara.exe",
		&sleepyCapybaraStarted,
		"-ConfigAddress="+config.BeastMaster.ConfigAddress,
	)
	for !sleepyCapybaraStarted {
	}

	go startModule(
		"./LazyRaven.exe",
		&lazyRavenStarted,
		"-ConfigAddress="+config.BeastMaster.ConfigAddress,
	)

	for !lazyRavenStarted {
	}
}

func startModule(path string, started *bool, args ...string) {
	for {
		scanner, cmd := internal.RunExecutableWithScanner(path, args...)

		*started = true
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		err := cmd.Wait()
		if err != nil {
			return
		}
		*started = false
		debug.Logf("%s module exited. Restarting module...", path)
	}

}

func RegisterRpcServices(config configuration.Config) error {
	configService := new(configuration.ConfigService)
	configService.Config = config

	err := rpc.Register(configService)
	if err != nil {
		return err
	}

	return nil
}
