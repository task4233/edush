package practice

import (
	"context"
	"github.com/docker/docker/client"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func Test_DocerPs(t *testing.T) {
	DockerPs()
}

func Test_DockerVersion(t *testing.T) {
	DockerVersion()
}

func Test_DockerRun(t *testing.T) {
	DockerRunSample()
}

func Test_Exec(t *testing.T) {
	type args struct {
		cmd []string
	}
	tests := map[string]struct {
		args args
	}{
		"ls":{
			args: args{
				cmd: []string{"ls"},
			},
		},
		"ls -a":{
			args: args{
				cmd: []string{"ls", "-a"},
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			ctx := context.Background()
			cli, _ := client.NewClientWithOpts(client.FromEnv)
			_, _ = Run(ctx, cli)
			if err := Exec(ctx, test.args.cmd, "test_container", cli); err != nil {
				panic(err)
			}
		})
	}
}
