package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// SendMessageFromBot - Отправка сообщения от лица бота
// token - токен бота
// chatID - id чата в который требуется отправить сообщение
func SendMessageFromBot(token, chatID, message string) (err error) {
	data := url.Values{
		"chat_id":    {chatID},
		"text":       {message},
		"parse_mode": {"MarkdownV2"},
	}

	var resp *http.Response
	resp, err = http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token), data)

	fmt.Println(&resp)

	return
}
