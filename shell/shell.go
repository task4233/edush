package shell

import (
	"fmt"
	"io/ioutil"
	"log"
	"bytes"
	"strings"

	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/taise-hub/edush/container"
	"github.com/taise-hub/edush/redis"
)


type ExecResult struct {
	Cmd    []byte
	StdOut []byte
	StdErr []byte
	Owner  bool
}

func CmdExecOnContainer(name string, p []byte) (ExecResult, error) {
	var execResult ExecResult
	var outBuf, errBuf bytes.Buffer
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println(err)
		return execResult, err
	}

	dir := GetCurrentDirecotry(name)
	reader, err := container.Exec(name, fmt.Sprintf("%s", p), dir, cli)
	if err != nil {
		log.Println(err)
		return execResult, err
	}
	outputDone := make(chan error)
	go func() {
		_, err := stdcopy.StdCopy(&outBuf, &errBuf, reader)
		outputDone <- err
	}()
	
	err = <-outputDone
	if err != nil {
		return execResult, err
	}

	stdout, err := ioutil.ReadAll(&outBuf)
	if err != nil {
		log.Println(err)
		return execResult, err
	}

	stderr, err := ioutil.ReadAll(&errBuf)
	if err != nil {
		log.Println(err)
		return execResult, err
	}
	i := strings.LastIndex(string(stdout), "\n")
	if i > -1 {
		stdout = stdout[:strings.LastIndex(string(stdout), "\n")]
	}
	j := strings.LastIndex(string(stdout), "\n")
	if j > -1 {
		pwd := string(stdout[j+1:])
		stdout = stdout[:j]
		SetCurrentDirectory(name, pwd)
		fmt.Printf("【DEBUG】pwd: %v\n", pwd)
	}

	execResult.StdOut = stdout
	execResult.StdErr = stderr
	return execResult, nil
}

func GetCurrentDirecotry(key string) string {
	client := redis.Connect(0)
	directory, err := client.Get(key).Result()
	if err != nil {
		log.Print(err)
		client.Set(key, "/", 0)
		return "/"
	}
	return directory
}

func SetCurrentDirectory(key, pwd string) {
	client := redis.Connect(0)
	client.Set(key, pwd, 0)
}
