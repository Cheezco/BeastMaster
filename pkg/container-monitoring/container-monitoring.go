package container_monitoring

import (
	"BeastMaster/internal/debug"
	"BeastMaster/pkg/configuration"
	"bytes"
	"context"
	"github.com/docker/docker/client"
	"log"
)

func MonitorContainer(containerId string, dataChan chan<- string, ctx context.Context, cli *client.Client) error {
	buf := new(bytes.Buffer)

	for {
		containerStats, err := cli.ContainerStats(ctx, containerId, false)
		if err != nil {
			continue
		}

		_, err = buf.ReadFrom(containerStats.Body)
		if err != nil {
			continue
		}

		dataChan <- buf.String()
		buf.Reset()
	}
}

func MonitorContainers(config configuration.LazyRaven, dataChan chan string, parsedDataChan chan ContainerData) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return err
	}

	for i := range config.Containers {
		index := i
		go func() {
			err := MonitorContainer(config.Containers[index].Id, dataChan, ctx, cli)
			if err != nil {
				debug.Log(err)
				log.Fatal(err)
			}
		}()
	}

	for i := 0; i < config.ParserCount; i++ {
		go func() {
			err := ParseData(dataChan, parsedDataChan)
			if err != nil {
				debug.Log(err)
				log.Fatal(err)
			}
		}()
	}

	return nil
}
