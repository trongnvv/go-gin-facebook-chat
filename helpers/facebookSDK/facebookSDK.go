package facebookSDK

import (
	"net/http"
	"os"
	"trongnv-chat/helpers"

	"github.com/gin-gonic/gin"
)

func GetPageInfo(c *gin.Context, pageID string) (map[string]interface{}, error) {
	res, err := helpers.FetchAPI(helpers.ParamsFetchAPI{
		GinContext: c,
		Url:        os.Getenv("API_ENDPOINT") + "/facebook/pages/" + pageID,
		Method:     http.MethodGet,
	})
	if err != nil {
		return nil, err
	}
	return res["data"].(map[string]interface{}), nil
}

func GetPageAccessToken(c *gin.Context, pageID string) (string, error) {
	res, err := helpers.FetchAPI(helpers.ParamsFetchAPI{
		GinContext: c,
		Url:        os.Getenv("API_ENDPOINT") + "/facebook/pages/" + pageID + "/token",
		Method:     http.MethodGet,
	})
	if err != nil {
		return "", err
	}
	return res["data"].(string), nil
}
