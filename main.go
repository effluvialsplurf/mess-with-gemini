package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func callModel(ctx context.Context, client *genai.Client, userInput string) {
	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(
		ctx,
		genai.Text(userInput),
	)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	fmt.Printf("\n---\n")
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get api key
	apiKey := os.Getenv("GOOGLE_AI_STUDIO_API_KEY")

	// get the context and our client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close() // close the client when we are done

	// get the user input and call the model
	scanner := bufio.NewScanner(os.Stdin)

	var userInput string
	fmt.Println("Please enter a prompt to be used by Gemini:")
	scanner.Scan()
	userInput = scanner.Text()
	fmt.Println("\nYou entered: ", userInput)
	callModel(ctx, client, userInput)
}
