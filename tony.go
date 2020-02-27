package tony

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

type Engine struct {
	context *gin.Context
	data    validate.DataFace
}

func New(c *gin.Context) *Engine {
	data, err := validate.FromRequest(c.Request)
	if err != nil {
		data = nil
	}
	return &Engine{c, data}
}
