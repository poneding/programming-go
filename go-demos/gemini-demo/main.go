package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const APIKEY = "GEMINI_API_KEY"

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(APIKEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	// chat
	for {
		// stdin reader
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		if input == "exit" {
			log.Println("Bye!")
			break
		}
		// generate response
		iter := model.GenerateContentStream(ctx, genai.Text(input))
		if err != nil {
			log.Fatal(err)
		}
		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			// print the generated text
			printResponse(resp)
		}
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, candidate := range resp.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				switch part.(type) {
				case genai.Text:
					log.Printf("%v", part)
				}
			}
		}
	}
}
