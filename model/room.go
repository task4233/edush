package model

type Room struct {
	ID string
	Clients []*Client 
	Forward chan []byte
	Join chan *Client
}

func NewRoom(id string) *Room {
	return &Room{
		ID: id,
		Clients: make([]*Client,0 , 2),
		Forward: make(chan []byte),
		Join: make(chan *Client),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <- r.Join:
			r.Clients = append(r.Clients, client)
		case msg := <- r.Forward:
			for _, client := range r.Clients {
				client.Message <- msg;
			}
		}
	}
}