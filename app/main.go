package main

import (
	"aws-sns-local-go/config"
	"aws-sns-local-go/internal/controller/rest"
	"aws-sns-local-go/internal/gateway/repository"
	"aws-sns-local-go/usecase/aws"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

const (
	defaultAddress = ":8080"
)

func init() {
	log.Printf("Loading .env file")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	address := os.Getenv("SERVER_ADDRESS")

	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Static("/", "views")

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Renderer = t

	db := config.DBConnect()

	topicRepo := repository.NewTopicRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	awsSvc := aws.NewService(topicRepo, messageRepo)
	rest.NewAwsHandler(e, awsSvc)

	e.GET("/health", healthCheck)

	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address))
}

func healthCheck(c echo.Context) error {
	return c.String(200, "OK")
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
