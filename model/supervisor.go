package model

import (
	"errors"
	"log"
	"github.com/gorilla/websocket"
)

//clientとroomの監督をする。
type Supervisor struct {
	clients map[string]*Client // "SESSID":*Client{}
	rooms map[string]*Room // "RoomName":*Room{}
}

func NewSupervisor() *Supervisor{
	return &Supervisor{
		clients: make(map[string]*Client),
		rooms: make(map[string]*Room),
	}
}

func (spv *Supervisor) Append(clientID, roomName string, conn *websocket.Conn)error {
	var room *Room
	var client *Client	
	
	if _, exist := spv.rooms[roomName]; !exist {
		room = NewRoom(roomName)
		spv.rooms[roomName] = room
		go room.run()
	}else {
		room = spv.rooms[roomName]
	}
	log.Println(room.Clients)
	
	if _, exist := spv.clients[clientID]; !exist {
		client = NewClient(clientID, conn, room)
		if success := client.joinRoom(); !success {
			return errors.New("I'm sorry, but I can only play this game with two people.")
		} 
		spv.clients[clientID] = client
		go client.read()
		log.Println("【DEBUG】run go client.read() done.")
		go client.write()
		log.Println("【DEBUG】run go client.write() done.")
		log.Println(client.Room)
	}
	log.Println("Supervisor.Append() done.")
	log.Println(spv.clients)
	log.Println(spv.rooms)
	return nil
}

/**
Supervisorの役割
- 外部から流入してくるメッセージを該当するroomに渡す。

Roomの役割
- 流入してきたメッセージをRoom内のClientに渡す。

- WebSocketを使って、送信する。
*/

/** 
【全体の流れ】
1. room作成
    → room名とclient名をpost
	→ Supervisorにroom名がなければ、Supervisorにclientとroomを追加する。
2. room参加
	→ room名とclient名をpost
	→ Supervisorにroom名があれば、clientだけ登録する。
	→ 既存のroomにclientを追加する。(もしclientが2名すでにいればエラー返す)
3. コマンド送信(clientはあるRoomに所属してることが前提)
	→ SupervisorがsessionIDを受け取る
	→ mapを使ってSessionIDからclientを識別する。
	→ client.Listen()でclient.HubにBroadcastする。

Supervisorに追加されているroomはRun()メソッドが走り、clientはListen(),Write()が動いている。
*/
