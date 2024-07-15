package http

import (
	"log"

	core "github.com/easyjoker/egame_core"
	"github.com/gin-gonic/gin"
)

// 接收 Player 的 json 資料，並將其新增至資料庫
func CreatePlayerHandler(c *gin.Context) {
	// 從請求中解析 Player 的 json 資料
	player := &core.Player{}
	if err := c.BindJSON(player); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// 創建一個新的 Player 實例
	// 將 Player 實例新增至資料庫
	// 回傳新增的 Player 實例

}
