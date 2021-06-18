package config

import "os"

var FACEBOOK_HOST string
var CHAT_BACKEND string
var MAIN_BACKEND string

func Init() {
	FACEBOOK_HOST = "https://graph.facebook.com"
	CHAT_BACKEND = os.Getenv("API_ENDPOINT_PUBLIC") + "/chat"
	MAIN_BACKEND = os.Getenv("API_ENDPOINT") + ""
}
