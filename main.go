package main

import (
	"egame_backend/server"
	"fmt"
	"log"
	"os"

	c "github.com/easyjoker/egame_core/config"
)

func main() {
	c.Init("config.yaml")

	config := c.GetConfig()
	log.SetOutput(os.Stdout)
	r := server.SetupRouter()

	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}
