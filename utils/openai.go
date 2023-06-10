package utils

import (
	"context"
	"fmt"
	"os"
	openai "github.com/sashabaranov/go-openai"
)

func Generator(prompt string) string {


	openaiKey := os.Getenv("OPENAI_KEY")
	c := openai.NewClient(openaiKey)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 1000,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return "Error"
	}
	return string(resp.Choices[0].Text)

}
