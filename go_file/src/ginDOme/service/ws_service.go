package service

import (
	"encoding/json"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Chan chan []byte //传递用户数据
	Addr string      //客户端的IP
}
type Msg struct {
	Content string // 消息内容
	User    string // 消息发布者
}

func HandleWebsocket(ws *websocket.Conn) {
	addr := ws.RemoteAddr().String()
	client := Client{make(chan []byte), addr}
	// content := client.Addr + "已上线!"
	// msg := Msg{Content: content, User: "系统消息"}
	// msgBytes, _ := json.Marshal(msg)
	// client.Chan <- msgBytes
	// fmt.Println(111)
	// a:=<-client.Chan
	go ReadMsgToClient(ws,client)
	go WriteMsgToClient(ws,client)
}
func WriteMsgToClient(ws *websocket.Conn, client Client) {
	for {
		msgBytes := <-client.Chan
		fmt.Println(msgBytes)
		err := ws.WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
func ReadMsgToClient(ws *websocket.Conn, client Client) {
	for {
		_, message, err := ws.ReadMessage()
		log.Println(message)
		if err != nil {
			log.Println("read error:", err)
			break
		}
		content := string(message)
		msg := Msg{Content: content, User: client.Addr}
		msgBytes, _ := json.Marshal(msg)
		client.Chan <- msgBytes

	}
}
