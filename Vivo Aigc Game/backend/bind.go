package main

import (
	"chatbox/route/ai"
	"chatbox/route/users"
	"chatbox/servicecontext"

	"github.com/gin-gonic/gin"
)

func Bind(server *gin.Engine, sctx *servicecontext.ServiceContext) {
	server.StaticFile("/test.html", "./test.html")

	server_ai := server.Group("/ai")
	server_ai.POST("/run", ai.NewRunRoute(sctx))

	server_users := server.Group("/user")
	server_users.POST("/register", users.NewRegisterRoute(sctx))
}
