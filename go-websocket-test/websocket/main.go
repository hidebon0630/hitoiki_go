package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// メッセージ用構造体
type Message struct {
	Msg  string `json:"msg"`
	Name string `json:"name"`
}

var clients = make(map[*websocket.Conn]bool) // 接続されるクライアント
var broadcast = make(chan Message)           // メッセージ用ブロードキャストチャネル

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketConnectHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	clients[conn] = true // 接続したクライアントを保存
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	msg.Msg = r.FormValue("msg")
	msg.Name = r.FormValue("name")
	broadcast <- msg // メッセージ用ブロードキャストチャネルに送信
}

func websocketMessage() {
	for {
		// チャネルからメッセージを受け取る
		msg := <-broadcast
		// 現在接続しているクライアントすべてにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err) // クライアントの接続が切れるとエラー
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	portNumber := "9000"
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", websocketConnectHandler)
	http.HandleFunc("/msg", messageHandler)
	log.Println("Server listening on port ", portNumber)
	go websocketMessage()
	http.ListenAndServe(":"+portNumber, nil)
}
