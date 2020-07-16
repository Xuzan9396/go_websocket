package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("http://test.room.wangran.live/ws/?port=9001", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn    *websocket.Conn
		err     error
		msgType int
		data    []byte
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	for {
		//text
		if msgType, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		_ = msgType

		fmt.Println(msgType, string(data))
		if err = conn.WriteMessage(websocket.TextMessage, []byte(string(data)+"xuzan")); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
