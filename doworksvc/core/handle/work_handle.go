package handle

import (
	"net/http"

	"doworksvc/pojo"
	"doworksvc/rsp"
	"doworksvc/util/v"
	"github.com/gin-gonic/gin"
)

// @Tags work
// @summary work新建
// @Accept json
// @Produce  json
// @Param text body pojo.Work true "work对象"
// @Router /api/work/add [POST]
// @Success 200 {object} rsp.Result
func WorkAdd(ctx *gin.Context) {
	param := new(pojo.Work)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	param.UserId = mustGetUser(ctx).Id
	data, err := svcFactory(ctx).WorkSvc().Add(param)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Data(data))
}

// @Tags work
// @summary work分页数据
// @Accept json
// @Produce  json
// @Param page query int false "当前页"
// @Param size query int false "每页条数"
// @Router /api/work/list [POST]
// @Success 200 {object} rsp.Result{data=rsp.Page{list=pojo.Work}}
func WorkList(ctx *gin.Context) {
	page, size := pageAndSize(ctx)
	list, total, err := svcFactory(ctx).WorkSvc().List(bindM(ctx), page, size)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.List(list, total))
}

// @Tags work
// @summary work删除
// @Accept json
// @Produce  json
// @Param id body v.Id true "要删除数据的id"
// @Router /api/work/del [post]
// @Success 200 {object} rsp.Result
func WorkDel(ctx *gin.Context) {
	param := new(v.Id)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).WorkSvc().Del(param.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}

type setText struct {
	Id   uint64 `json:"id" example:"1"`
	Text string `json:"text" example:"修改的内容"`
}

// @Tags work
// @summary work修改内容
// @Accept json
// @Produce  json
// @Param id body setText true "要修改数据的id"
// @Router /api/work/setText [post]
// @Success 200 {object} rsp.Result
func WorkSetText(ctx *gin.Context) {
	param := new(pojo.Work)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).WorkSvc().SetText(param)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}

type setCompleted struct {
	Id        uint64 `json:"id" example:"1"`
	Completed bool   `json:"completed" example:"true"`
}

// @Tags work
// @summary work修改状态
// @Accept json
// @Produce  json
// @Param id body setCompleted true "要修改数据的id"
// @Router /api/work/setCompleted [post]
// @Success 200 {object} rsp.Result
func WorkSetCompleted(ctx *gin.Context) {
	param := new(pojo.Work)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).WorkSvc().SetCompleted(param.Id, param.Completed)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}

type setBelongDate struct {
	Id         uint64 `json:"id" example:"1"`
	BelongDate string `json:"belong_date" example:"2006-01-02"`
}

// @Tags work
// @summary work修改所属日期
// @Accept json
// @Produce  json
// @Param id body setBelongDate true "要修改数据的id"
// @Router /api/work/setBelongDate [post]
// @Success 200 {object} rsp.Result
func WorkSetBelongDate(ctx *gin.Context) {
	param := new(pojo.Work)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).WorkSvc().SetBelongDate(param.Id, param.BelongDate)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}
