package http

import (
	"log"
	"strconv"

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

	result, err := p.NewPlayer(player.Account, player.Password, player.Name, 1000)
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

	result, err := p.LoginPlayer(player.Account, player.Password)

	if err != nil {
		log.Printf("Login error: %v", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, result)
}

// add player money
func AddMoneyHandler(c *gin.Context) {
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		log.Printf("ParseInt error: %v", parseErr)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	money, amountErr := strconv.ParseFloat(c.PostForm("amount"), 64)

	if amountErr != nil {
		log.Printf("ParseFloat error: %v", amountErr)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	err := p.AddPlayerBalance(id, money)

	if err != nil {
		log.Printf("AddMoney error: %v", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
