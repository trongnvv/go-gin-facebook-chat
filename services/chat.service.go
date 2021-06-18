package services

import (
	"errors"
	"net/http"
	"net/url"
	"trongnv-chat/config"
	"trongnv-chat/define"
	"trongnv-chat/helpers"
)

type ChatService struct{}

func (s ChatService) ConvertPageMessages(paging interface{}, pageID string, conversationID string) (interface{}, error) {
	if paging == nil {
		return nil, errors.New("paging new")
	}
	rs := map[string]interface{}{}
	page := paging.(map[string]interface{})
	if page["next"] != nil {
		param, err := url.ParseQuery(page["next"].(string))
		if err != nil {
			return nil, err
		}
		rs["next"] = config.CHAT_BACKEND + "/" + pageID + "/messages?conversationID=" + conversationID + "&after=" + param.Get("after")
	}

	if page["previous"] != nil {
		param, err := url.ParseQuery(page["previous"].(string))
		if err != nil {
			return nil, err
		}
		rs["previous"] = config.CHAT_BACKEND + "/" + pageID + "/messages?conversationID" + conversationID + "&before=" + param.Get("before")
	}
	return rs, nil
}

func (s ChatService) ConvertPageConversations(paging interface{}, pageID string) (interface{}, error) {
	if paging == nil {
		return nil, errors.New("paging new")
	}
	rs := map[string]interface{}{}
	page := paging.(map[string]interface{})
	if page["next"] != nil {
		param, err := url.ParseQuery(page["next"].(string))
		if err != nil {
			return nil, err
		}
		rs["next"] = config.CHAT_BACKEND + "/" + pageID + "/conversations?after=" + param.Get("after")
	}

	if page["previous"] != nil {
		param, err := url.ParseQuery(page["previous"].(string))
		if err != nil {
			return nil, err
		}
		rs["previous"] = config.CHAT_BACKEND + "/" + pageID + "/conversations?before=" + param.Get("before")
	}
	return rs, nil
}

func (s ChatService) FetchMessages(params define.ParamsMessages) (map[string]interface{}, error) {
	res, err := helpers.FetchAPI(helpers.ParamsFetchAPI{
		Url:    config.FACEBOOK_HOST + "/" + params.ConversationID + "/messages",
		Method: http.MethodGet,
		Query: struct {
			AccessToken string `url:"access_token"`
			Fields      string `url:"fields,omitempty"`
			Limit       int    `url:"limit,omitempty"`
			Before      string `url:"before,omitempty"`
			After       string `url:"after,omitempty"`
		}{
			AccessToken: params.AccessToken,
			Fields:      "created_time,id,from,message",
			Limit:       1,
			Before:      params.Before,
			After:       params.After,
		},
	})
	return res, err
}

func (s ChatService) FetchConversations(params define.ParamsConversations) (map[string]interface{}, error) {
	res, err := helpers.FetchAPI(helpers.ParamsFetchAPI{
		Url:    config.FACEBOOK_HOST + "/" + params.PageFBID + "/conversations",
		Method: http.MethodGet,
		Query: struct {
			AccessToken string `url:"access_token"`
			Fields      string `url:"fields,omitempty"`
			Limit       int    `url:"limit,omitempty"`
			Before      string `url:"before,omitempty"`
			After       string `url:"after,omitempty"`
		}{
			AccessToken: params.AccessToken,
			Fields:      "unread_count,snippet,message_count,link,updated_time,participants",
			Limit:       1,
			Before:      params.Before,
			After:       params.After,
		},
	})
	return res, err
}
