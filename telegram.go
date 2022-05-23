package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// SendMessageFromBot - Отправка сообщения от telegram бота
// token - токен бота
// chatID - id чата в который требуется отправить сообщение
func SendMessageFromBot(token, chatID, message string) (err error) {
	data := url.Values{
		"chat_id": {chatID},
		"text":    {message},
	}

	var resp *http.Response
	if resp, err = http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token), data); err != nil {
		return
	}

	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != 200 {
		var respData []byte
		if respData, err = io.ReadAll(resp.Body); err != nil {
			return
		}

		return fmt.Errorf("ошибка при отправке сообщения в telegram: %d - %s", resp.StatusCode, string(respData))
	}

	return
}
