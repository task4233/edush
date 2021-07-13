package model

type CmdQueue struct {
	Pipe chan []byte
}

func NewCmdQueue() *CmdQueue {
	return &CmdQueue{
		Pipe: make(chan []byte),
	}
}