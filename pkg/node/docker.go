package node

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerContainer struct {
}

func CreateClient() {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.43"))
	if err != nil {
		fmt.Println(err)
	}
	containers, err := client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println(err)
	}
	for _, containerIns := range containers {
		fmt.Printf("%+v", containerIns)
		data, err := client.ContainerInspect(context.TODO(), containerIns.ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("pid:", data.State.Pid)
		fmt.Println("port:", data.NetworkSettings.Ports)
	}
}
