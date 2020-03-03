package tony

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/tidwall/gjson"
)

func GetInt(c *gin.Context, key string, defaultValue ...int) int {
	r := getData(c, key)
	res, err := strconv.Atoi(fmt.Sprintf("%v", r))
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func GetFloat(c *gin.Context, key string, defaultValue ...float64) float64 {
	r := getData(c, key)
	res, err := strconv.ParseFloat(fmt.Sprintf("%v", r), 64)
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func GetString(c *gin.Context, key string, defaultValue ...string) string {
	r := getData(c, key)
	res := fmt.Sprintf("%v", r)
	if res == "<nil>" {
		res = ""
	}
	if res == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return res
}

func GetSlice(c *gin.Context, key string) []interface{} {
	r := getData(c, key)
	res, _ := r.([]interface{})
	return res
}

func GetMap(c *gin.Context, key string) map[string]interface{} {
	r := getData(c, key)
	res, _ := r.(map[string]interface{})
	return res
}

func GetGJSON(c *gin.Context, key string) *gjson.Result {
	r := getData(c, key)
	res := map[string]interface{}{
		key: r,
	}
	bt, err := json.Marshal(res)
	if err != nil {
		return nil
	}
	result := gjson.ParseBytes(bt).Get(key)
	return &result
}

func getData(c *gin.Context, key string) interface{} {
	rdata, ok := c.Get("requestData")
	if !ok {
		data, err := validate.FromRequest(c.Request)
		c.Set("requestData", data)
		if err != nil {
			return nil
		}
		res, _ := data.Get(key)
		return res
	}
	res, _ := rdata.(validate.DataFace).Get(key)
	return res
}
