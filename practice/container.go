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
	time.Sleep(20 * time.Second)
	timeout := 30 * time.Second
	if err := cli.ContainerStop(context.Background(), "test_container", &timeout); err != nil {
		panic(err)
	}
}


//コンテナを起動する。
func DockerRun(cli *client.Client) (string, error) {
	bg := context.Background()
	cc := &container.Config{
		Image:        "nginx",
		ExposedPorts: nat.PortSet{nat.Port("80"): struct{}{}},
	}
	hc := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("80"): []nat.PortBinding{nat.PortBinding{HostPort: "8080"}},
		},
		AutoRemove: true,
	}
	body, err := cli.ContainerCreate(bg, cc, hc, nil, nil, "test_container")
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(bg, body.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	return body.ID, nil
}
/**課題
ユーザに紐つける必要があるもの
- コンテナのID
- カレントディレクトリ
**/

//マンドを引数にとってコンテナ内でそのコマンドを実行させる方法例
/**
id: コンテナID DockerRun()の戻り値とか
**/
func DockerExec(id string, cli *client.Client) error {
	
	ec := &types.ExecConfig{
		User: "hoge",
		Privileged: false,
		Tty: true,
		WorkingDir: "/",
		Cmd: []string{"ls", "-a"},
	}
	idResponse, err := cli.ContainerExecCreate(context.Background(), id, *ec)
	if err != nil {
		return err
	}
	execStartCheck := types.ExecStartCheck{
		Detach: true,
		Tty:    false,
	}
	if err = cli.ContainerExecStart(context.Background(), idResponse.ID, execStartCheck); err != nil {
		return err
	}
	return nil
}
