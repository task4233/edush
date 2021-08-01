package model

type Room struct {
	ID string
	Clients []*Client 
	Forward chan []byte
}

func NewRoom(id string) *Room {
	return &Room{
		ID: id,
		Clients: make([]*Client,0 , 2),
		Forward: make(chan []byte),
	}
}

func (r *Room) run() {
	for {
		msg:= <- r.Forward
		for _, client := range r.Clients {
			client.Message <- msg
		}
	}
}