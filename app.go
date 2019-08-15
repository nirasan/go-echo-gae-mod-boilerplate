package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/nirasan/go-echo-gae-mod-boilerplate/app"
)

func main() {
	e := echo.New()

	e.Renderer = app.NewTemplate("templates/*.html")

	e.GET("/", app.IndexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Printf("Defaulting to port %s", port)
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
