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
	"github.com/taise-hub/edush/model"
	"github.com/taise-hub/edush/redis"
)

func CmdExecOnContainer(name string, p []byte) (model.ExecResult, error) {
	var execResult model.ExecResult
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println(err)
		return execResult, err
	}
	dir := GetCurrentDirecotry(name)
	log.Print(fmt.Sprintf("%s", p))
	reader, err := container.Exec(name, fmt.Sprintf("%s", p), dir, cli)
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

func ChangeDirectory(key ,newdir string) error {

	client := redis.Connect(0)
	olddir, err := client.Get(key).Result()
	if err != nil {
		log.Print(err)
		return err
	}
	switch {
	case newdir[0:1] == "./":
		directory := olddir + newdir[1:]
		return client.Set(key, directory, 0).Err()
	case newdir[0:2] == "../":
		index := strings.Index(Reverse(olddir), "/")
		directory := olddir[:len(olddir)-index-1] + newdir[2:]
		return client.Set(key, directory, 0).Err()
	}
	return nil
}

func GetCurrentDirecotry(key string) string {
	client := redis.Connect(0)
	directory, err := client.Get(key).Result()
	if err != nil {
		log.Print(err)
		return "/"
	}
	return directory
}

func Reverse(s string) string {

    runes := []rune(s)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}
