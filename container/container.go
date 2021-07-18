package container

import (
	"context"
	// "fmt"
	// "time"
	// "log"
	// "bufio"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	// "github.com/docker/go-connections/nat"
	// "github.com/docker/docker/pkg/stdcopy"
)

func Run(name string, cli *client.Client) error {
	cc := &container.Config{
		Image:        "nginx",
	}
	hc := &container.HostConfig{
		AutoRemove: true,
	}
	body, err := cli.ContainerCreate(context.Background(), cc, hc, nil, nil, name)
	if err != nil {
		return err
	}
	if err := cli.ContainerStart(context.Background(), body.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}