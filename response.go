package tony

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResSuccess(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		data[0] = ""
	}

	c.JSON(200, Response{
		Code: 0,
		Msg:  "success",
		Data: data[0],
	})
}

func ResError(c *gin.Context, code int, err interface{}) {
	msg := ""
	if e, ok := err.(error); ok {
		msg = e.Error()
	} else {
		msg = fmt.Sprintf("%v", err)
	}

	c.AbortWithStatusJSON(200, Response{
		Code: code,
		Msg:  msg,
	})
}
