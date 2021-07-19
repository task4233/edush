package shell

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/docker/docker/client"
	"github.com/taise-hub/edush/container"
)


func CmdExecOnContainer(name string, p []byte) ([]byte, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	reader, err := container.Exec(name, fmt.Sprintf("%s", p), cli)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//TODO: stdout, stderr両方返すようにする。
	output, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("%v", output)
	return output, nil
}
