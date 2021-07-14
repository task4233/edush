package practice

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// $ docker ps
func DockerPs() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID, container.Image)
	}
}

func DockerVersion() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	v := cli.ClientVersion()
	fmt.Printf("%s\n", v)
}

func DockerRun() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	//imageの設定
	cc := &container.Config{
		Image:        "nginx",
		ExposedPorts: nat.PortSet{nat.Port("80"): struct{}{}},
	}

	//ホストの設定
	hc := &container.HostConfig{
		// --port 8080:80
		PortBindings: nat.PortMap{
			nat.Port("80"): []nat.PortBinding{nat.PortBinding{HostPort: "8080"}},
		},
		// --rm
		AutoRemove: true,
	}
	body, err := cli.ContainerCreate(context.Background(), cc, hc, nil, nil, "test_container")
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerStart(context.Background(), body.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	//起動してます。
	time.Sleep(20 * time.Second)
	//終了処理に入ります。
	timeout := 30 * time.Second
	if err := cli.ContainerStop(context.Background(), "test_container", &timeout); err != nil {
		panic(err)
	}
}
