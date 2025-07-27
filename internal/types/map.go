package types

func (tu *TelegramUpdate) AsBindInput() *BindInput {
	return &BindInput{
		TelegramID: tu.Message.From.ID,
		ChatID:     tu.Message.Chat.ID,
		ChatTitle:  tu.Message.Chat.Title,
	}
}
