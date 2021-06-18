package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
)

type ParamsFetchAPI struct {
	GinContext *gin.Context
	Url        string
	Method     string
	Query      interface{}
	Body       interface{}
}

func FetchAPI(params ParamsFetchAPI) (map[string]interface{}, error) {
	body := make([]byte, 0)
	if params.Body != nil {
		body, _ = json.Marshal(params.Body)
	}
	v := url.Values{}
	if params.Query != nil {
		v, _ = query.Values(params.Query)
		params.Url = params.Url + "?" + v.Encode()
	}

	req, err := http.NewRequest(params.Method, params.Url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if params.GinContext != nil {
		if authorization := params.GinContext.GetHeader("Authorization"); authorization != "" {
			req.Header.Set("Authorization", authorization)
		}

		if tracerID := params.GinContext.GetHeader("Authorization"); tracerID != "" {
			req.Header.Set("uber-trace-id", tracerID)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rs map[string]interface{}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &rs)
	if resp.StatusCode == 200 {
		return rs, nil
	}
	textErr := "Fetch API error!"
	// if rs["message"] != nil {
	// 	textErr = rs["message"].(string)
	// } else if rs["errors"] != nil {
	// 	textErr = rs["errors"].(string)
	// } else if rs["error"] != nil {
	// 	textErr = rs["error"].(string)
	// }
	log.Println("FetchAPI error: status", resp.StatusCode, rs)
	return nil, errors.New(textErr)
}
