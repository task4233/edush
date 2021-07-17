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
	type args struct {
		name string
	}
	tests := map[string]struct {
		args args
	}{
		"コンテナの作成例": {
			args: args {
				name: "test",
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			ctx := context.Background()
			cli, _ := client.NewClientWithOpts(client.FromEnv)
			
			if err := Run(ctx, test.args.name, cli); err != nil {
				panic(err)
			}
		})
	}
}

func Test_Exec(t *testing.T) {
	type args struct {
		cmd []string
		name string
	}
	tests := map[string]struct {
		args args
	}{
		"ls":{
			args: args{
				cmd: []string{"ls"},
				name: "test1",
			},
		},
		"ls -a":{
			args: args{
				cmd: []string{"ls", "-a"},
				name: "test2",
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			ctx := context.Background()
			cli, _ := client.NewClientWithOpts(client.FromEnv)
			_ = Run(ctx, test.args.name, cli)
			
			if _, err := Exec(ctx, test.args.cmd, test.args.name, cli); err != nil {
				panic(err)
			}
		})
	}
}
