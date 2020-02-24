package tony

import "github.com/gin-gonic/gin"

type Engine struct {
	context *gin.Context
}

func New(c *gin.Context) *Engine {
	return &Engine{c}
}
