package main

import (
	"user-service/internal/databases"
	"user-service/internal/server"
)

func main() {
	databases.ConnectDB()
	databases.CreateAdmin()
	server.Start()
}
