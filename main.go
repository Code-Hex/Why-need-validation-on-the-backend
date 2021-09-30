package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	indexPage, err := template.New("index.html").Parse(simplePage)
	if err != nil {
		return errors.WithStack(err)
	}

	e := echo.New()

	e.Renderer = &Template{templates: indexPage}

	e.GET("/", Index())
	e.POST("/vote", Vote())

	return e.Start(":1323")
}

var storage = map[string]int{
	"banana": 0,
	"orange": 0,
}

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Title":  "Banana or Orange",
			"Fruits": storage,
		})
	}
}

func Vote() echo.HandlerFunc {
	return func(c echo.Context) error {
		fruit := c.FormValue("fruit")
		storage[fruit]++
		return Index()(c)
	}
}

const simplePage = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
	<ul>
{{ range $fruit, $count := .Fruits }}
      <li>
	    <span>{{ $fruit }}: {{ $count }}</span>
	  </li>
{{ end }}
    </ul>
    <form action="/vote" method="post">
      <label for="fruit">Which do you like?</label>
      <input type="text" id="fruit" name="fruit" pattern="banana|orange" required />
	  <button type="submit">vote</button>
    </form>
  </body>
</html>`

type Template struct {
	templates *template.Template
}

var _ echo.Renderer = (*Template)(nil)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
