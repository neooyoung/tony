package tony

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

func (e *Engine) GetInt(key string, defaultValue ...int) int {
	r := e.getData(key)
	res, err := strconv.Atoi(fmt.Sprintf("%v", r))
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func (e *Engine) GetFloat(key string, defaultValue ...float64) float64 {
	r := e.getData(key)
	res, err := strconv.ParseFloat(fmt.Sprintf("%v", r), 64)
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func (e *Engine) GetString(key string, defaultValue ...string) string {
	r := e.getData(key)
	res, ok := r.(string)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func (e *Engine) GetSlice(key string) []interface{} {
	r := e.getData(key)
	res, _ := r.([]interface{})
	return res
}

func (e *Engine) GetMap(key string) map[string]interface{} {
	r := e.getData(key)
	res, _ := r.(map[string]interface{})
	return res
}

func (e *Engine) GetGJSON(key string) *gjson.Result {
	r := e.getData(key)
	res := map[string]interface{}{
		key: r,
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bt, err := json.Marshal(res)
	if err != nil {
		return nil
	}
	result := gjson.ParseBytes(bt).Get(key)
	return &result
}

func (e *Engine) getData(key string) interface{} {
	if e.data == nil {
		return nil
	}
	r, _ := e.data.Get(key)
	return r
}
