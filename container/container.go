package container

import (
	"bufio"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Run(name string, cli *client.Client) error {
	cc := &container.Config{
		Image: "nginx",
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

func Exec(name string, cmd string, cli *client.Client) (*bufio.Reader, error) {

	cmds := []string{"/bin/bash", "-c", cmd}
	ec := &types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   "/",
		Cmd:          cmds,
	}
	idResp, err := cli.ContainerExecCreate(context.Background(), name, *ec)
	if err != nil {
		return nil, err
	}

	resp, err := cli.ContainerExecAttach(context.Background(), idResp.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, err
	}
	return resp.Reader, nil
}
