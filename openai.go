package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getReviewFromOpenAI(requirement string, reviewContent string) []string {

	// Get the review from OpenAI
	// https://beta.openai.com/docs/api-reference/answers/generate
	client := &http.Client{}
	prompt := `
		Just review the diff below.
		Below is requirement for this code:
		` + requirement + `
		Please check requirements and review the code below.
		` + reviewContent

	messages := []map[string]interface{}{
		{
			"role":    "user",
			"content": prompt,
		},
		{
			"role": "assistant",
			"content": `
				You are a pro engineer. Please review the code.
				Please answer in English.
				No need break down the code, just focus on the quality of the code.
				No need to write a long review.
				Please answer in 2-3 sentences.
			`,
		},
	}

	body := map[string]interface{}{
		"model":       "gpt-4o-mini",
		"messages":    messages,
		"max_tokens":  1000,
		"temperature": 0.5,
	}
	jsonBody, err := json.Marshal(body)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	openaiKey := os.Getenv("OPENAI_API_KEY")
	organization := os.Getenv("OPENAI_ORGANIZATION_ID")
	project := os.Getenv("OPENAI_PROJECT")
	req.Header.Set("Authorization", "Bearer "+openaiKey)
	req.Header.Set("OpenAI-Organization", organization)
	req.Header.Set("OpenAI-Project", project)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the response body
	var review map[string]interface{}
	err = json.Unmarshal(reqBody, &review)
	if err != nil {
		panic(err)
	}

	fmt.Println(review["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string))

	return []string{}
}
