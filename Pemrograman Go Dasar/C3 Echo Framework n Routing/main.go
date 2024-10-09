package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	// r.GET("/index", func(ctx echo.Context) error{
	// 	data := "Hello from /index"
	// 	return ctx.String(http.StatusOK, data)
	// })

	r.GET("/html", func(ctx echo.Context) error {
		data := "Hello from /html"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/index", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/html")
	})

	r.GET("/json", func(ctx echo.Context) error {
		data := M{"message": "Hello from /json", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	//Parsing request'
	//parsing query string
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	//parsing url path param
	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})
	r.Start(":9000")

	//parsing url path param dan setelahnya
	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	//Parsing form data
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	})
}
