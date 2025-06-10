package main

import (
	"chatbox/config"
	"chatbox/servicecontext"

	"github.com/gin-gonic/gin"
)

func main() {
	//给出APP_ID和APPKey，并生成ID
	sctx := servicecontext.NewServiceContext(config.Load())

	//服务器
	server := gin.New()
	//route
	Bind(server, sctx)
	// bind := server.Group("/ai")
	// bind.GET("/run", ai.NewRunRoute(&sctx))
	e := server.Run("127.0.0.1:5500")
	if e != nil {
		panic(e)
	}
	// message := []vivo.ChatMessage{
	// 	{
	// 		Role:    vivo.CHAT_ROLE_USER,
	// 		Content: "\n",
	// 	},
	// }
	// res, e := app.Chat(vivo.GenerateRequestID(), session_id, message, nil)
	// if e != nil {
	// 	panic(e)
	// }
	// message = append(message, res)
	// for i := 1; i <= 3; i++ {
	// 	_, e = fmt.Scan(&content)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// 	if content[0] == '1' {
	// 		break
	// 	}
	// 	message = append(message, vivo.ChatMessage{
	// 		Role:    vivo.CHAT_ROLE_USER,
	// 		Content: content,
	// 	})
	// 	res, e = app.Chat(vivo.GenerateRequestID(), session_id, message, nil)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// 	fmt.Println(res.Content)
	// 	message = append(message, res)
	// }
}
