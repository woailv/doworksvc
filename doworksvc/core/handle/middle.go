package handle

import (
	"doworksvc/core/svc"
	"doworksvc/db"
	"doworksvc/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"xorm.io/xorm"
)

const svcFactoryKey = "svcFactoryKey"
const dbSessionKey = "dbSessionKey"

//注入服务工厂
func InjectSvcFactory(ctx *gin.Context) {
	session := db.DB().NewSession()
	ctx.Set(dbSessionKey, session)
	f := svc.NewSvcFactory(session)
	ctx.Set(svcFactoryKey, f)
	ctx.Next()
	_ = session.Close()
}

func svcFactory(ctx *gin.Context) *svc.SvcFactory {
	return ctx.MustGet(svcFactoryKey).(*svc.SvcFactory)
}

func dbSession(ctx *gin.Context) *xorm.Session {
	return ctx.MustGet(dbSessionKey).(*xorm.Session)
}

//登录校验
func Auth(ctx *gin.Context) {
	_, ok := getUser(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	ctx.Next()
}

//查询限定到用户数据
func LimitUserData(ctx *gin.Context) {
	dbSession(ctx).Where("user_id = ?", mustGetUser(ctx).Id)
}

//会话管理
var store = sessions.NewCookieStore([]byte("sk"))

func startSession(ctx *gin.Context, user *pojo.User) {
	session, err := store.Get(ctx.Request, "user_key")
	if err != nil {
		panic(fmt.Sprintf("开启session异常:%s", err))
	}
	session.Values["user"] = user
	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		panic(fmt.Sprintf("保存session异常:%s", err))
	}
}

func endSession(ctx *gin.Context) {
	session, err := store.Get(ctx.Request, "user_key")
	if err != nil {
		panic(fmt.Sprintf("获取session异常"))
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		panic(fmt.Sprintf("结束session失败,%s", err))
	}
}

func getUser(ctx *gin.Context) (*pojo.User, bool) {
	session, err := store.Get(ctx.Request, "user_key")
	if err != nil {
		panic(fmt.Sprintf("开启session异常:%s", err))
	}
	user, ok := session.Values["user"].(*pojo.User)
	if !ok {
		return nil, false
	}
	return user, true
}

func mustGetUser(ctx *gin.Context) *pojo.User {
	user, ok := getUser(ctx)
	if !ok {
		panic("获取登录用户失败")
	}
	return user
}
