package types

type TelegramUpdate struct {
	Message struct {
		From struct {
			ID int64 `json:"id"`
		} `json:"from"`
		Chat struct {
			ID    int64  `json:"id"`
			Title string `json:"title"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}

type TelegramUpdateReply struct {
	Method      string      `json:"method"`
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}
