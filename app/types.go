package main

import (
  "context"
)

type WebAppInfo struct {
  URL string `json:"url"`
}

type InlineKeyboardButton struct {
  Text string      `json:"text"`
  WebApp *WebAppInfo `json:"web_app,omitempty"`
}

type ReplyMarkup struct {
  InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type Message struct {
  Text string      `json:"text"`
  ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type BindInput struct {
	TelegramID int64 `json:"telegram_id"`
	ChatID     int64 `json:"chat_id"`
	ChatTitle  string `json:"chat_title"`
}

type Service interface {
  Start(context.Context) Message 
  Bind(context.Context, BindInput) Message 
}

type TelegramUpdate struct {
  Message struct {
      From struct {
          ID int64 `json:"id"`
      } `json:"from"`
      Chat struct {
          ID    int64 `json:"id"`
          Title string `json:"title"`
      } `json:"chat"`
      Text string `json:"text"`
  } `json:"message"`
}

type TelegramUpdateReply struct{
  Method      string      `json:"method"`
  ChatID int64 `json:"chat_id"`
  Text string `json:"text"`
  ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}