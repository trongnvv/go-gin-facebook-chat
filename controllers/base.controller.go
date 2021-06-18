package controllers

import (
	"net/http"
	"time"
	"trongnv-chat/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseController struct{}

func (u BaseController) defaultInsertDB() *models.BaseSchema {
	return &models.BaseSchema{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u BaseController) responseSuccess(c *gin.Context, data interface{}, message string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": true, "data": data, "message": message})
}

func (u BaseController) responseError(c *gin.Context, status int, err string, message string) {
	c.AbortWithStatusJSON(status, gin.H{"success": false, "message": message, "error": err})
}
