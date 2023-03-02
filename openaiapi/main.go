package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

const ASAK = "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"

type Payload struct {
	Prompt           string  `json:"prompt"`
	MaxTokens        int64   `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             int64   `json:"top_p"`
	FrequencyPenalty int64   `json:"frequency_penalty"`
	PresencePenalty  int64   `json:"presence_penalty"`
	Model            string  `json:"model"`
}

func main() {

	data := Payload{
		Prompt:           "原神",
		MaxTokens:        2048,
		Temperature:      0.5,
		TopP:             0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Model:            "text-davinci-003",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", ASAK))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	text := jsoniter.Get(all, "choices", 0, "text").ToString()
	fmt.Println(text)
}
