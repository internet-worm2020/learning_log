package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	// "github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Chan chan []byte     // 传递用户数据
	Conn *websocket.Conn // 客户端的IP
	UID  uint
}
type Msg struct {
	Content  string `json:"content"` // 消息内容
	User     string `json:"user"`
	Reliever uint   `json:"reliever"`
	Sender   uint   `json:"sender"`
}

var clients = make(map[*websocket.Conn]*Client)
var broadcast = make(chan Msg)

func HandleWebsocket(ws *websocket.Conn, c *gin.Context) {
	uIDInterface, exists := c.Get("uId")
	if !exists {
		return
	}
	uID, ok := uIDInterface.(uint)
	if !ok {
		return
	}
	client := &Client{Chan: make(chan []byte), Conn: ws, UID: uID}
	clients[ws] = client
	go ReadMsgToClient(client)
	go WriteMsgToClient(client)
}
func WriteMsgToClient(client *Client) {
	for {
		msgBytes := <-client.Chan
		err := client.Conn.WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
func ReadMsgToClient(client *Client) {
	var msg Msg
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println("JSON unmarshaling error:", err)
			return
		}
		broadcast <- msg
	}
}

func MessageManager() {
	for {
		msg := <-broadcast
		for _, client := range clients {
			if client.UID == msg.Sender || client.UID == msg.Reliever {
				msgBytes, _ := json.Marshal(msg)
				client.Chan <- msgBytes
			}
		}
	}
}
func Acc(client *Client) {
	defer client.Conn.Close()
	defer close(client.Chan)
	defer delete(clients, client.Conn)
	for {
		_, _, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// 连接正常关闭或正在关闭
				return
			}
			// 处理其他错误情况
			return
		}
	}

}
