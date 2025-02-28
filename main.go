package main

import (
	"os"

	"github.com/harryrford/locust/chat"
	"github.com/harryrford/locust/web"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client := chat.NewClient(&chat.Config{
		Model:    os.Getenv("MODEL"),
		APIKey:   os.Getenv("API_KEY"),
		Endpoint: os.Getenv("MODEL_ENDPOINT"),
	})

	webClient := web.NewClient(client)

	webClient.Research("psychedlic mushrooms")
}
