package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

const (
	PoligonPath     = "https://api.polygon.io"
	TickerPath      = PoligonPath + "/v3/reference/tickers"
	DailyValuesPath = PoligonPath + "/v1/open-close"
)

var (
	Port   string
	ApiKey string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ApiKey = os.Getenv("API_KEY")
	Port = os.Getenv("PORT")
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err)
			return nil
		},
	})

	app.Get("/", HandleIndex)
	app.Get("/search", HandleSearch)
	app.Get("/values/:ticker", HandleValues)

	log.Fatal(app.Listen(":" + Port))
}
