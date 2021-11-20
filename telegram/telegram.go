package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	parseMode = "MarkdownV2"
)

type Telegram struct {
	token           string
	communicationId string
}

type Message struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func CreateTelegram(token, communicationId string) Telegram {
	return Telegram{token: token, communicationId: communicationId}
}

func (t *Telegram) SendMessage(message *string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.token)
	messageToSend := Message{ChatId: string(t.communicationId), Text: *message, ParseMode: parseMode}
	postBody, err := json.Marshal(&messageToSend)
	if err != nil {
		log.Fatalln("Cannot serialize telegram object")
	}
	log.Println(string(postBody))

	requestBody := bytes.NewBuffer(postBody)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatalf("An error occured %v", err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

}
