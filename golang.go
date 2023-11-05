package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/1147159208207327312/DrzCCCxuo_hiLZDpJPNJ2YUvzCQRfiHeBdCcLIcmhQ2zt0kO9O8U9ECuMV71lUV8Zfzu"

	data := map[string]string{
		"content": "Привет, Discord! Это сообщение отправлено через вебхук.",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при маршалинге JSON:", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Сообщение успешно отправлено!")
	} else {
		fmt.Println("Произошла ошибка при отправке сообщения. Код:", resp.StatusCode)
	}
}
