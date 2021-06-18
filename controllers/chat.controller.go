package controllers

import (
	"net/http"
	"trongnv-chat/define"
	"trongnv-chat/helpers/facebookSDK"
	"trongnv-chat/services"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	*BaseController
	service *services.ChatService
}

func NewChatController() *ChatController {
	controller := new(BaseController)
	service := new(services.ChatService)
	return &ChatController{BaseController: controller, service: service}
}

func (u ChatController) GetMessages(c *gin.Context) {
	pageID := c.Param("pageID")
	before := c.Query("before")
	after := c.Query("after")
	conversationID := c.Query("conversationID")
	respond := map[string]interface{}{}
	accessToken, err := facebookSDK.GetPageAccessToken(c, pageID)
	if err != nil {
		u.responseError(c, http.StatusInternalServerError, err.Error(), "")
		return
	}
	res, err1 := u.service.FetchMessages(define.ParamsMessages{
		AccessToken:    accessToken,
		ConversationID: conversationID,
		Before:         before,
		After:          after,
	})

	if err1 != nil {
		u.responseError(c, http.StatusInternalServerError, err1.Error(), "")
		return
	}

	paging, _ := u.service.ConvertPageMessages(res["paging"], pageID, conversationID)
	results := []interface{}{}
	for _, v := range res["data"].([]interface{}) {
		_v := v.(map[string]interface{})
		from := _v["from"].(map[string]interface{})
		results = append(results, map[string]interface{}{
			"message":     _v["message"],
			"createdTime": _v["created_time"],
			"name":        from["name"],
		})
	}

	respond["paging"] = paging
	respond["results"] = results
	u.responseSuccess(c, respond, "")
	return
}

func (u ChatController) GetConversations(c *gin.Context) {
	pageID := c.Param("pageID")
	before := c.Query("before")
	after := c.Query("after")
	respond := map[string]interface{}{}

	pageFB, err := facebookSDK.GetPageInfo(c, pageID)
	if err != nil {
		u.responseError(c, http.StatusInternalServerError, err.Error(), "")
		return
	}

	res, err1 := u.service.FetchConversations(define.ParamsConversations{
		AccessToken: pageFB["accessToken"].(string),
		PageFBID:    pageFB["pageFacebookID"].(string),
		Before:      before,
		After:       after,
	})

	if err1 != nil {
		u.responseError(c, http.StatusInternalServerError, err1.Error(), "")
		return
	}

	paging, _ := u.service.ConvertPageConversations(res["paging"], pageID)
	results := []interface{}{}
	for _, v := range res["data"].([]interface{}) {
		_v := v.(map[string]interface{})
		results = append(results, map[string]interface{}{
			"snippet":        _v["snippet"],
			"unreadCount":    _v["unread_count"],
			"messageCount":   _v["message_count"],
			"link":           "https://facebook.com/" + _v["link"].(string),
			"updatedTime":    _v["updated_time"],
			"conversationID": _v["id"],
			"participants":   _v["participants"].(map[string]interface{})["data"],
		})
	}
	respond["paging"] = paging
	respond["results"] = results
	u.responseSuccess(c, respond, "")
}
