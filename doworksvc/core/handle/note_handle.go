package handle

import (
	"net/http"

	"doworksvc/pojo"
	"doworksvc/rsp"
	"doworksvc/util/v"
	"github.com/gin-gonic/gin"
)

// @Tags 笔记
// @summary note新建
// @Accept json
// @Produce  json
// @Param text body pojo.Note true "note对象"
// @Router /api/note/add [POST]
// @Success 200 {object} rsp.Result
func NoteAdd(ctx *gin.Context) {
	param := new(pojo.Note)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	data, err := svcFactory(ctx).NoteSvc().Add(param)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Data(data))
}

// @Tags 笔记
// @summary note分页数据
// @Accept json
// @Produce  json
// @Param page query int false "当前页"
// @Param size query int false "每页条数"
// @Router /api/note/list [POST]
// @Success 200 {object} rsp.Result{data=rsp.Page{list=pojo.Note}}
func NoteList(ctx *gin.Context) {
	list, total, err := svcFactory(ctx).NoteSvc().List(pageAndSize(ctx))
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.List(list, total))
}

// @Tags 笔记
// @summary note删除
// @Accept json
// @Produce  json
// @Param id body v.Id true "要删除数据的id"
// @Router /api/note/del [post]
// @Success 200 {object} rsp.Result
func NoteDel(ctx *gin.Context) {
	param := new(v.Id)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).NoteSvc().Del(param.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}

// @Tags 笔记
// @summary note更新
// @Accept json
// @Produce  json
// @Param param body pojo.Note true "更新成的对象"
// @Router /api/note/update [post]
// @Success 200 {object} rsp.Result
func NoteUpdate(ctx *gin.Context) {
	param := new(pojo.Note)
	if err := ctx.BindJSON(param); err != nil {
		ctx.JSON(http.StatusBadRequest, rsp.Err(err))
		return
	}
	err := svcFactory(ctx).NoteSvc().Update(param)
	if err != nil {
		ctx.JSON(http.StatusOK, rsp.Err(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp.Ok())
}
