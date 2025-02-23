package main

import (
	"fmt"
	"log"
	"os"

	"github.com/harryrford/locust/chat"
	"github.com/harryrford/locust/locust"
	"github.com/joho/godotenv"
)

func main() {

	// sx, sy := robotgo.GetScreenSize()
	// bit := robotgo.CaptureScreen(0, 0, sx, sy)
	// defer robotgo.FreeBitmap(bit)

	// img := robotgo.ToImage(bit)
	// imgo.Save("test.png", img)

	err := godotenv.Load()
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}

	client := chat.NewClient(&chat.Config{
		Model:    os.Getenv("MODEL"),
		APIKey:   os.Getenv("API_KEY"),
		Endpoint: os.Getenv("MODEL_ENDPOINT"),
	})

	resp, err := locust.DeepResearch(client, "Build a blog website with Django and React")
	if err != nil {
		panic(err)
	}
	// resp, err := locust.GetFinalAnswer(client, "What is the capital of France?", []locust.SubquestionAnswer{
	// 	{
	// 		Subquestion: "What is the capital of France?",
	// 		Answer:      "Paris",
	// 		Sources:     []string{"https://en.wikipedia.org/wiki/Paris"},
	// 	},
	// 	{
	// 		Subquestion: "What is the population of Paris?",
	// 		Answer:      "2,148,271",
	// 		Sources:     []string{"https://en.wikipedia.org/wiki/Paris"},
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(resp)
	// fmt.Println("test")
	// fmt.Println(resp, err)

	// app := fiber.New()

	// app.Get("/research/:query", func(c fiber.Ctx) error {
	// 	resposne, err := locust.DeepResearch(client, c.Params("query"))
	// 	if err != nil {
	// 		return c.SendStatus(http.StatusInternalServerError)
	// 	}

	// 	return c.SendString(resposne)
	// })

	// log.Fatal(app.Listen(":3000"))
}
