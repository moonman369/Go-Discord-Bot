package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moonman369/Go-Discord-Bot/config"
)

type MessageTemplate struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int `json:"index"`
	Message      MessageTemplate
	FinishReason string `json:"finish_reason"`
}

type UsageTemplate struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Response struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []Choice
	Usage   UsageTemplate
}

func SendPrompt(promptContent string) Response {
	url := "https://api.openai.com/v1/chat/completions"
	reqBody := []byte(fmt.Sprintf(`{
    "model": "gpt-3.5-turbo",
    "messages": [
        {
            "role": "user",
            "content": "%v"
        }
    ],
    "temperature": 0.7,
    "max_tokens": 500
}`, promptContent))
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.OpenAIKey))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		panic(err)
	}

	// fmt.Println(resp.Choices[0].Message.Content)

	return resp
}
