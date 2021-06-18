package helpers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := FetchAPI(ParamsFetchAPI{
			GinContext: c,
			Url:        os.Getenv("API_ENDPOINT") + "/check-auth",
			Method:     http.MethodGet,
		})
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "Unauthorized",
				})
			return
		}
    c.Set("user", res["data"])
		c.Next()
	}
}
