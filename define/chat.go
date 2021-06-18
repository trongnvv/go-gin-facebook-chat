package define

type ParamsConversations struct {
	AccessToken, PageFBID, Before, After string
}

type ParamsMessages struct {
	AccessToken, ConversationID, Before, After string
}
