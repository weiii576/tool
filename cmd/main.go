package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/api"
	"github.com/weiii576/tool/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	server := gin.Default()

	api.SetRouter(server, store)

	server.Run()
}
