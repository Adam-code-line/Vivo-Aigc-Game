package ai

import (
	"chatbox/servicecontext"

	"github.com/dingdinglz/vivo"
	"github.com/gin-gonic/gin"
)

type Question struct {
	Context string `json:"question"`
}

func NewRunRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var question Question
		err := ctx.ShouldBindJSON(&question)
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "error",
			})
			panic((err))
		}
		sctx.Message = append(sctx.Message, vivo.ChatMessage{
			Role:    vivo.CHAT_ROLE_USER,
			Content: question.Context,
		})
		res, e := sctx.Config.App.Chat(vivo.GenerateRequestID(), sctx.Config.Session, sctx.Message, nil)
		if e != nil {
			panic(e)
		}
		sctx.Message = append(sctx.Message, res)
		ctx.JSON(200, gin.H{
			"message": res.Content,
		})
	}
}
