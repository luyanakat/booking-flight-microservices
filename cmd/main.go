package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	api "mock-project/cmd/server"
)

func main() {
	fmt.Println("Listening on port: 3000")
	if err := api.Server().Run(":3000"); err != nil {
		log.Println("Get error from run server", zap.Error(err))
	}
}
