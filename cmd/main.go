package main

import (
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type data struct {
	color uint8
	pos   uint8
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "static")

	t := NewTemplates()
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	list := make([]data, 100)
	for i := uint8(0); i < 100; i++ {
		list[i].color = uint8(rand.Intn(4))
		list[i].pos = i
	}
	message := fmt.Sprintf("%v", list)
	e.GET("/data", func(c echo.Context) error {
		return c.String(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
