package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/harryrford/locust/chat"
	"github.com/harryrford/locust/locust"
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

	app := fiber.New()

	app.Get("/research/:query", func(c fiber.Ctx) error {
		resposne, err := locust.DeepResearch(client, c.Params("query"))
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.SendString(resposne)
	})

	log.Fatal(app.Listen(":3000"))
}
