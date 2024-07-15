package main

import (
	"egamebackend/server"
	"fmt"
	"log"

	"egamebackend/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	r := server.SetupRouter()

	r.Run(fmt.Sprintf(":%s", config.Port))
}
