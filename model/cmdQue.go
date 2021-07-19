package model

type CmdQueue struct {
	ResultPipe chan ExecResult
}

func NewCmdQueue() *CmdQueue {
	return &CmdQueue{
		ResultPipe: make(chan ExecResult),
	}
}
