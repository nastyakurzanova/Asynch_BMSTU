package main

import (
	"fmt"
	webapi "lab8"
	"lab8/internal/pkg/handler"
	"log"
)

func main() {
	handlers := handler.NewHandler()

	server := new(webapi.Server)
	fmt.Println("Server is running on:", webapi.GetOutboundIP(), ":", "8070")

	if err := server.Run("8070", handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to init server, error: %v", err)
	}

}
