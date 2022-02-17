package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/server"
)

func main() {

	db.Init()

	server := server.NewServer()
	server.Run()
}
