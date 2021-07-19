package shell

import (
	"fmt"
	"io/ioutil"
	"log"
	"bytes"

	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/taise-hub/edush/container"
	"github.com/taise-hub/edush/model"
)

func CmdExecOnContainer(name string, p []byte) (model.ExecResult, error) {
	var execResult model.ExecResult
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println(err)
		return execResult, err
	}
	reader, err := container.Exec(name, fmt.Sprintf("%s", p), cli)
	if err != nil {
		log.Println(err)
		return execResult, err
	}
	var outBuf, errBuf bytes.Buffer
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

	fmt.Printf("stdout: %v\n", string(stdout))
	fmt.Printf("stderr: %v\n", string(stderr))
	execResult.StdOut = stdout
	execResult.StdErr = stderr

	return execResult, nil
}
