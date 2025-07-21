package main

func (tu *TelegramUpdate) asBindInput() *BindInput {
	return &BindInput{
		TelegramID: tu.Message.From.ID,
		ChatID:     tu.Message.Chat.ID,
		ChatTitle:  tu.Message.Chat.Title,
	}
}
