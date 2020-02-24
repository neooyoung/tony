package tony

import "fmt"

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (e *Engine) ResSuccess(data interface{}) {
	e.context.JSON(200, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func (e *Engine) ResError(code int, err interface{}) {
	msg := ""
	if e, ok := err.(error); ok {
		msg = e.Error()
	} else {
		msg = fmt.Sprintf("%v", err)
	}

	e.context.AbortWithStatusJSON(200, Response{
		Code: code,
		Msg:  msg,
	})
}