package parse

import (
	"img_tag/pkg/consts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context) *Gin {
	return &Gin{C: c}
}

type Gin struct {
	C *gin.Context
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func (g *Gin) Res200(data interface{}) {
	g.Response(http.StatusOK, consts.SUCCESS, "", data)
}

func (g *Gin) Res400(errCode int, msg string, data interface{}) {
	g.Response(http.StatusBadRequest, errCode, msg, data)
}

func (g *Gin) Response(httpCode, errCode int, msg string, data interface{}) {
	if ok, _ := strconv.ParseBool(msg); !(ok) {
		msg = consts.ERROR_MAP[errCode]
	}
	g.C.JSON(
		httpCode,
		response{
			Code: errCode,
			Data: data,
			Msg:  msg,
		},
	)
}
