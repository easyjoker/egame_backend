package ws

import (
	"encoding/json"
	"log"

	"github.com/easyjoker/pokers/lobby"
)

// 處理用戶送上來的命令
func HandleGameCmd(playerId int64, channel *chan string, cmd_type *string, cmd interface{}) {
	// 將命令轉發給遊戲服務器
	l := lobby.GetHome()
	c := CheckCmd(cmd)
	if c == nil {
		log.Printf("Invalid command: %v", cmd)
		return
	}

	l.ReceiveCommand(playerId, channel, cmd_type, c)
}

// 檢查命令是否合法
func CheckCmd(cmd interface{}) *lobby.LobbyCommand {
	if cmd == nil {
		return nil
	}

	j, err := json.Marshal(cmd)
	if err != nil {
		log.Printf("Could not marshal command: %v", err)
		return nil
	} else {
		var r lobby.LobbyCommand = lobby.LobbyCommand{}
		e := json.Unmarshal(j, &r)
		if e != nil {
			log.Printf("Could not unmarshal command: %v", e)
			return nil
		}

		return &r
	}
}
