package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/goccy/go-json"
)
type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func main() {
	// Replace with your OpenAI API key
	apiKey := os.Getenv("OPENAI_API_KEY") // Store your API key in an environment variable

	// Define the request payload
	requestBody := OpenAIRequest{
		Model: "gpt-4o-mini", // Use "gpt-3.5-turbo" if you want a cheaper option
		Messages: []Message{
			{"system", "You are a helpful assistant."},
			{"user", "give me a list of best restaurants in the Newyork city"},
		},
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create HTTP request
	url := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse response
	var openAIResponse OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print response
	if len(openAIResponse.Choices) > 0 {
		fmt.Println("OpenAI Response:", openAIResponse.Choices[0].Message.Content)
	} else {
		fmt.Println("No response received from OpenAI")
	}
}
