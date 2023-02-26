package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	godotenv.Load()
	apikey := os.Getenv("API_KEY")

	ctx := context.Background()
	client := gpt3.NewClient(apikey)

	root := &cobra.Command{
		Use:   "Chat GPT",
		Short: "Chat with chat gpt",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false
			for !quit {
				fmt.Print("Say something (send 'quit' to exit): ")
				if !scanner.Scan() {
					break
				}

				ques := scanner.Text()
				switch ques {
				case "quit":
					quit = true

				default:
					GetResponse(client, ctx, ques)
				}
			}
		},
	}
	root.Execute()
}

func GetResponse(client gpt3.Client, ctx context.Context, ques string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:      []string{ques},
		MaxTokens:   gpt3.IntPtr(40),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		log.Println(err)
		os.Exit(13)
	}
	fmt.Printf("\n")
}
