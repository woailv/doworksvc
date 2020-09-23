package handle

import (
	"doworksvc/util/v"
	"strconv"

	"github.com/gin-gonic/gin"
)

func pageAndSize(ctx *gin.Context) (int, int) {
	page := ctx.DefaultQuery("page", "1")
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	size := ctx.DefaultQuery("size", "10")
	s, err := strconv.Atoi(size)
	if err != nil {
		s = 1
	}
	return p, s
}

func bindM(ctx *gin.Context) v.M {
	m := v.M{}
	_ = ctx.BindJSON(&m)
	return m
}
