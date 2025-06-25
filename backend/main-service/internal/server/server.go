package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	s := gin.Default()

	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
