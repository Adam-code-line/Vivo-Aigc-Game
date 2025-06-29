package users

import (
	"chatbox/servicecontext"
	"chatbox/tool"
	"chatbox/types"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Telephone string `json:"telephone" bg:"must"`
	Password  string `json:"password" bg:"must"`
}
type RegisterResp struct {
	types.BaseResp
}

func NewRegisterRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	fmt.Println("111")
	return func(ctx *gin.Context) {
		req := &RegisterReq{}
		e := tool.ParseForm(req, ctx)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		resp, e := register(sctx, req)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		ctx.JSON(200, resp)
	}
}

func register(sctx *servicecontext.ServiceContext, req *RegisterReq) (resp *RegisterResp, e error) {

	// 此处可以做对telephone的校验：
	// 比如是否为手机号的正确格式
	// 增加验证码逻辑防止恶意注册

	// 尝试创建一个新用户
	// Password设为tool.GenerateMD5(req.Password)的原因是防止数据库泄漏

	// e = sctx.UserModel.Create(&database.UserModel{
	// 	Telephone: req.Telephone,
	// 	Password:  tool.GenerateMD5(req.Password),
	// })

	// 这里对没发生错误做判断的原因是，如果有错误，e包含了错误内容，resp可以为nil
	// 没有错误才需要放置成功提示

	// if e == nil {
	// 	resp = &RegisterResp{
	// 		BaseResp: types.BaseResp{
	// 			Code:    0,
	// 			Message: "注册成功！",
	// 		},
	// 	}
	// }
	return
}
