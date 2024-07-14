package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithCode(code int, c *gin.Context) {
	msg, ok := c2m[code]
	if !ok {
		msg = "unknown error"
	}
	Result(code, map[string]any{}, msg, c)
}
