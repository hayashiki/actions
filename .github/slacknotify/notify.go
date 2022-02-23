package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type payload struct {
	Text string `json:"text"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type string `json:"type"`
	Text *Text  `json:"text,omitempty"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func notify(webhookURL, message string) (err error) {
	p, err := json.Marshal(payload{
		Text: message,
		Blocks: []Block{{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: message,
			},
		}},
	})
	if err != nil {
		return err
	}
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
