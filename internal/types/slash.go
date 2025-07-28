package types

type Message struct {
	Text        string       `json:"text"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup"`
}

type BindInput struct {
	TelegramID int64  `json:"telegram_id"`
	ChatID     int64  `json:"chat_id"`
	ChatTitle  string `json:"chat_title"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}
