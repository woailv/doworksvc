package handle

import (
	"doworksvc/pojo"
	"doworksvc/rsp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags 用户
// @summary 注册
// @Accept json
// @Produce  json
// @Param param body pojo.UserRegisterParam true "用户注册参数"
// @Router /api/user/register [POST]
// @Success 200 {object} rsp.Result
func UserRegister(ctx *gin.Context) {
	param := new(pojo.UserRegisterParam)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	data, err := svcFactory(ctx).UserSvc().Register(param)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Data(data))
}

// @Tags 用户
// @summary 登录
// @Accept json
// @Produce  json
// @Param param body pojo.UserLoginParam true "用户登录参数"
// @Router /api/user/login [POST]
// @Success 200 {object} rsp.Result
func UserLogin(ctx *gin.Context) {
	param := new(pojo.UserLoginParam)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	data, err := svcFactory(ctx).UserSvc().Login(param)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	startSession(ctx, data)
	ctx.JSON(http.StatusOK, rsp.Data(data))
}

// @Tags 用户
// @summary 退出登录
// @Router /api/user/logout [GET]
// @Success 200 {object} rsp.Result
func UserLogout(ctx *gin.Context) {
	endSession(ctx)
	ctx.JSON(http.StatusOK, rsp.Ok())
}

type userRegisterEmailKey struct {
	Email string `json:"email" binding:"required"`
}

// @Tags 用户
// @summary 获取注册邮箱验证码
// @Accept json
// @Produce  json
// @Param param body userRegisterEmailKey true "邮箱验证码参数"
// @Router /api/user/registerEmailKey [POST]
// @Success 200 {object} rsp.Result
func UserRegisterEmailCode(ctx *gin.Context) {
	param := new(userRegisterEmailKey)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).UserSvc().UserRegisterEmailCode(param.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}
