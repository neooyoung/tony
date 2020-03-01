package tony

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	jsoniter "github.com/json-iterator/go"
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

func GetContentType(req *http.Request) interface{} {
	contentTypeHeader := req.Header.Get("Content-Type")
	contentTypeSlice := strings.Split(contentTypeHeader, ";")
	fmt.Println(contentTypeSlice)
	var res interface{}
	switch contentTypeSlice[0] {
	case "":
		res = parseQuery(req)
	case "application/json":
		res = parseJSON(req)
	case "multipart/form-data":
		res = parseForm(req)
	case "application/x-www-form-urlencoded":
		res = parseWWWForm(req)
	default:
		res = nil
	}
	return res
}

func parseQuery(req *http.Request) interface{} {
	query := req.URL.Query()
	return parseUrlValues(&query)
}

func parseJSON(req *http.Request) interface{} {
	data := getBodyData(req)
	var res map[string]interface{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(data, &res)
	return res
}

func parseForm(req *http.Request) interface{} {
	err := req.ParseMultipartForm(1024)
	if err != nil {
		return nil
	}
	return parseUrlValues(&req.PostForm)
}

func parseWWWForm(req *http.Request) interface{} {
	data := getBodyData(req)
	urlValues, err := url.ParseQuery(string(data))
	if err != nil {
		return nil
	}
	return parseUrlValues(&urlValues)
}

func getBodyData(req *http.Request) []byte {
	body := req.Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}
	return data
}

func parseUrlValues(urlValues *url.Values) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range *urlValues {
		if len(v) == 1 {
			res[k] = v[0]
		} else {
			res[k] = v
		}
	}
	return res
}
