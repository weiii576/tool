package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/api"
	"github.com/weiii576/tool/configs"
	"github.com/weiii576/tool/storage"
)

func main() {
	env := configs.NewEnv()

	store, err := storage.NewPostgresStore(env)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	server := gin.Default()

	api.SetRouter(server, store, env)

	server.Run()
}
