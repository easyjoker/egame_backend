package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]*Connection)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	// 註冊新的客戶端並記錄當前的 Unix 時間戳
	clients[conn] = NewConnection(conn)
	log.Println("New connection")
	// 啟動一個新的 goroutine 來處理這個連接
	go func() {
		for {
			defer conn.Close()
			// if c := clients[conn]; c != nil {
			// 	// 超過 60 秒沒有收到訊息就斷開連接
			// 	if time.Now().Unix() > c.LastMessageTime+60000 {
			// 		msg := fmt.Sprintf("[%d]:Connection [%s] closed due to inactivity", c.PlayerId, conn.RemoteAddr().Network())
			// 		log.Println(msg)
			// 		delete(clients, conn)
			// 		break
			// 	}
			// }
			// message type: 1 = text, 2 = binary, 3 = close, 4 = ping, 5 = pong
			messageType, p, err := conn.ReadMessage()
			log.Printf("messageType: %d\n", messageType)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Unexpected close error: %v", err)
				} else {
					log.Printf("Read error: %v", err)
				}
				delete(clients, conn)
				break
			}

			if messageType == 3 {
				log.Printf("收到關閉信息\n")
				delete(clients, conn)
				break
			}

			if messageType == 1 {
				// 將訊息轉成 message
				message := Message[interface{}]{}
				err := json.Unmarshal(p, &message)
				if err != nil {
					log.Printf("Could not unmarshal message: %v\n", err)
					continue
				}

				if message.Type == "register" {
					if message.Data == nil {
						log.Printf("Could not register connection: no player id provided\n")
						continue
					}
					// 註冊玩家Id
					playerId, ok := message.Data.(float64)
					if !ok {
						log.Printf("Could not register connection: invalid player id\n")
						continue
					}
					clients[conn].SavePlayerId(int64(playerId))
					log.Printf("[%d]:Connection [%s] registered\n", int64(playerId), conn.RemoteAddr().Network())
					conn.WriteJSON(Message[interface{}]{Type: "register", Data: "success"})
					continue
				}
				if clients[conn].PlayerId == 0 {
					log.Printf("Could not process message: connection not registered\n")
					conn.WriteJSON(Message[interface{}]{Type: "error", Data: "not registered"})

					delete(clients, conn)
					break
				}
				log.Printf("收到信息： %v,將資料往後送到遊戲處理去\n", message)
				conn.WriteJSON(Message[interface{}]{Type: "received", Data: "command received"})
				c := clients[conn]
				HandleGameCmd(c.PlayerId, &c.Channel, &message.Type, message.Data)
			}
		}
	}()

	// 啟動一個新的 goroutine 來發送信息到用戶端
	go func() {
		for {
			select {
			case message := <-clients[conn].Channel:
				log.Printf("發送信息給用戶端：%v\n", message)
				err := conn.WriteJSON(message)
				if err != nil {
					log.Printf("Could not send message: %v\n", err)
					return
				}
			}
		}
	}()
}
