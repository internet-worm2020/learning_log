package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gindome/pkg"
)

type ResponseData struct {
	Code pkg.ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code pkg.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code pkg.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: pkg.CodeSuccess,
		Msg:  pkg.CodeSuccess.Msg(),
		Data: data,
	})
}
func ResponseOperateSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: pkg.CodeSuccess,
		Msg:  pkg.CodeSuccess.Msg(),
		Data: nil,
	})
}
