package http

import (
	"log"

	p "github.com/easyjoker/egame_core/player"
	"github.com/gin-gonic/gin"
)

// 接收 Player 的 json 資料，並將其新增至資料庫
func CreatePlayerHandler(c *gin.Context) {
	// 從請求中解析 Player 的 json 資料
	player := &p.Player{}
	if err := c.BindJSON(player); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// 檢查參數是否合法
	if player.Account == "" || player.Password == "" || player.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	result, err := p.NewPlayer(nil, player.Account, player.Password, player.Name, 1000)
	if err != nil {
		log.Printf("NewPlayer error: %v", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, result)
}

// 用戶登入
func LoginHandler(c *gin.Context) {
	// 從請求中解析 Player 的 json 資料
	player := &p.Player{}
	if err := c.BindJSON(player); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// 檢查參數是否合法
	if player.Account == "" || player.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	result, err := p.LoginPlayer(nil, player.Account, player.Password)

	if err != nil {
		log.Printf("Login error: %v", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, result)
}
