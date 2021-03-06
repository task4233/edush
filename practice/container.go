package practice

import (
	"context"
	"fmt"
	"time"
	"log"
	"bufio"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	// "github.com/docker/docker/pkg/stdcopy"
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

// dockerのバージョン確認例
func DockerVersion() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	v := cli.ClientVersion()
	fmt.Printf("%s\n", v)
}

//コンテナを起動する例
func DockerRunSample() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cc := &container.Config{
		Image:        "nginx",
		ExposedPorts: nat.PortSet{nat.Port("80"): struct{}{}},
	}

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
	time.Sleep(20 * time.Second)
	timeout := 30 * time.Second
	if err := cli.ContainerStop(context.Background(), "test_container", &timeout); err != nil {
		panic(err)
	}
}


//コンテナを起動する。
func Run(ctx context.Context,name string, cli *client.Client) error {
	cc := &container.Config{
		Image:        "nginx",
	}
	hc := &container.HostConfig{
		AutoRemove: true,
	}
	body, err := cli.ContainerCreate(ctx, cc, hc, nil, nil, name)
	if err != nil {
		return err
	}
	if err := cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}
/**
id: コンテナID 名前でもok
**/
func Exec(ctx context.Context, cmd []string, id string, cli *client.Client) (*bufio.Reader, error) {
	
	ec := &types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir: "/",
		Cmd: cmd,
	}
	
	idResp, err := cli.ContainerExecCreate(ctx, id, *ec)
	if err != nil {
		return nil, err
	}

	resp, err := cli.ContainerExecAttach(ctx, idResp.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Conn.Close(); err != nil {
			log.Panic(err)
		}
		log.Println("connection closed")
	}()

	return resp.Reader, nil
}


/**課題
ユーザに紐つける必要があるもの
- コンテナのID→(userID+乱数)でコンテナ立てるようにする。
- カレントディレクトリ→TODO
**/
