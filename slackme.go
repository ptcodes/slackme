package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func format(hostname, message string) string {
	currentTime := time.Now().Format("15:05:05")
	return fmt.Sprintf("[%s %s] %s", currentTime, hostname, message)
}

func send(slackWebhookURL, text string) error {
	jsonMessage, err := json.Marshal(SlackMessage{Text: text})
	if err != nil {
		return err
	}

	resp, err := http.Post(slackWebhookURL, "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(body) != "ok" {
		return fmt.Errorf(string(body))
	}
	return nil
}

func main() {
	slackWebhookURL := os.Getenv("SLACKME_WEBHOOK_URL")
	if slackWebhookURL == "" {
		log.Fatal("error: SLACKME_WEBHOOK_URL cannot be empty")
	}

	if len(os.Args) < 2 {
		log.Fatal("usage: ./slackme <your message>")
	}
	message := strings.Join(os.Args[1:], " ")

	if len(message) == 0 {
		log.Fatal("error: please provide a message")
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "no hostname"
	}

	formattedMessage := format(hostname, message)

	err = send(slackWebhookURL, formattedMessage)
	if err != nil {
		log.Fatal(err)
	}
}
