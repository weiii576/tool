package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weiii576/tool/api"
)

func main() {
	server := gin.Default()

	api.SetRouter(server)

	server.Run()
}
