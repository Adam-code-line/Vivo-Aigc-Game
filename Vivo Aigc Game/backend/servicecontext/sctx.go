package servicecontext

import (
	"chatbox/config"

	"github.com/dingdinglz/vivo"
)

type ServiceContext struct {
	Config  *config.Config
	Message []vivo.ChatMessage
	// UserModel *database.UserModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	message := []vivo.ChatMessage{
		{
			Role:    vivo.CHAT_ROLE_USER,
			Content: "你的名字是慧编，是一个帮助我学习编程的助手，可以回答我关于编程的问题，会给予我解决问题的思路，但不要直接发送代码。\n",
		},
	}
	res, e := c.App.Chat(vivo.GenerateRequestID(), c.Session, message, nil)
	if e != nil {
		panic(e)
	}
	message = append(message, res)
	return &ServiceContext{
		Config:  c,
		Message: message,
		// UserModel: database.NewUserModel(),
	}
}
