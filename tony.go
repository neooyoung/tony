package tony

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

func GetDatas(key string, req *http.Request) interface{} {
	data, err := validate.FromRequest(req)
	if err != nil {
		data = nil
	}
	// res, _ := data.Get(key)
	return data
}

func GetRequestData(req *http.Request) (map[string]interface{}, error) {
	contentType := getContentType(req)
	res := make(map[string]interface{})
	var err error

	switch contentType {
	case "":
		res = parseQuery(req)
	case "application/json":
		res, err = parseJSON(req)
	case "multipart/form-data":
		res, err = parseForm(req)
	case "application/x-www-form-urlencoded":
		res, err = parseWWWForm(req)
	}

	return res, err
}

func getContentType(req *http.Request) string {
	contentTypeHeader := req.Header.Get("Content-Type")
	contentTypeSlice := strings.Split(contentTypeHeader, ";")
	return contentTypeSlice[0]
}

func parseQuery(req *http.Request) map[string]interface{} {
	query := req.URL.Query()
	return parseURLValues(&query)
}

func parseJSON(req *http.Request) (map[string]interface{}, error) {
	data, err := getBodyData(req)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	// json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(data, &res)
	return res, nil
}

func parseForm(req *http.Request) (map[string]interface{}, error) {
	return nil, nil
	// err := req.ParseMultipartForm(1024)
	// if err != nil {
	// 	return nil, err
	// }

	// return parseURLValues(&req.MultipartForm.Value), nil
}

func parseWWWForm(req *http.Request) (map[string]interface{}, error) {
	err := req.ParseForm()
	if err != nil {
		return nil, err
	}

	// data, err := getBodyData(req)
	// if err != nil {
	// 	return nil, err
	// }

	// urlValues, err := url.ParseQuery(string(data))
	// if err != nil {
	// 	return nil, err
	// }
	return parseURLValues(&req.PostForm), nil
}

func getBodyData(req *http.Request) ([]byte, error) {
	body := req.Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseURLValues(urlValues *url.Values) map[string]interface{} {
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
