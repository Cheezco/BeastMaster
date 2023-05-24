package container_monitoring

import (
	"BeastMaster/internal"
	"BeastMaster/internal/debug"
	"BeastMaster/pkg/configuration"
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"log"
	"strconv"
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

	debug.Log("Spreading docker compose")

	containers := spreadComposeContainers(config.Containers, ctx, cli)

	debug.Log("Finished spreading docker compose")
	debug.Log("Container count: ", len(containers))

	for i := range containers {
		index := i
		go func() {
			err := MonitorContainer(containers[index].Id, dataChan, ctx, cli)
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

func spreadComposeContainers(containers []configuration.Container, ctx context.Context, cli *client.Client) []configuration.Container {
	result := make([]configuration.Container, 0)
	for _, container := range containers {
		if container.Compose == "" {
			debug.Log("Compose not used " + container.Id)
			result = append(result, container)
			continue
		}
		debug.Log("Compose used " + container.Compose)

		//var args = filters.Args{}
		//args.Add("label", )
		flt := filters.NewArgs()
		flt.Add("label", "com.docker.compose.project="+container.Compose)

		cnt, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: flt})

		debug.Log("Found " + strconv.Itoa(len(cnt)) + " containers")
		internal.CheckError(err)
		for _, composeContainer := range cnt {
			debug.Log("Appending " + composeContainer.ID)
			result = append(result, configuration.Container{Id: composeContainer.ID})
		}
	}

	return result
}
