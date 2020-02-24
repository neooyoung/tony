package tony

import (
	"fmt"
	"strconv"

	"github.com/gookit/validate"
)

func (e *Engine) GetInt(key string, defaultValue ...int) int {
	r := e.GetData(key)
	res, err := strconv.Atoi(fmt.Sprintf("%v", r))
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func (e *Engine) GetFloat(key string, defaultValue ...float64) float64 {
	r := e.GetData(key)
	res, err := strconv.ParseFloat(fmt.Sprintf("%v", r), 64)
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func (e *Engine) GetString(key string, defaultValue ...string) string {
	r := e.GetData(key)
	res, ok := r.(string)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func (e *Engine) GetSlice(key string) []interface{} {
	r := e.GetData(key)
	res, _ := r.([]interface{})
	return res
}

func (e *Engine) GetMap(key string) map[string]interface{} {
	r := e.GetData(key)
	res, _ := r.(map[string]interface{})
	return res
}

func (e *Engine) GetData(key string) interface{} {
	data, err := validate.FromRequest(e.context.Request)
	if err != nil {
		return nil
	}
	res, _ := data.Get(key)
	return res
}
