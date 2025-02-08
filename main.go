package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/harryrford/locust/locust"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	app := fiber.New()

	app.Get("/grok/:question", func(c fiber.Ctx) error {

		return c.SendString("Thinking:\n" + locust.NewLocustQuery(apiKey, c.Params("question")))
	})

	log.Fatal(app.Listen(":3000"))
}
