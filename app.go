package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "hello world")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Printf("Defaulting to port %s", port)
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
