package ws

import (
	"time"

	"github.com/gorilla/websocket"
)

// Connection 代表一個websocket連接
type Connection struct {
	// 這個連接的websocket
	Conn *websocket.Conn
	// 這個連接的玩家Id
	PlayerId uint64
	// 最後一次收到訊息的時間
	LastMessageTime int64
}

// NewConnection 用來創建一個新的連接
func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		Conn:            conn,
		LastMessageTime: time.Now().Unix(),
	}
}

func (c *Connection) SavePlayerId(playerId uint64) {
	c.PlayerId = playerId
	c.LastMessageTime = time.Now().Unix()
}

// SendMessage 用來發送訊息給這個連接
func (c *Connection) SendMessage(messageType int, data []byte) error {
	return c.Conn.WriteMessage(messageType, data)
}

// SendMessageString 用來發送字串訊息給這個連接
func (c *Connection) SendMessageString(data string) error {
	return c.SendMessage(websocket.TextMessage, []byte(data))
}
