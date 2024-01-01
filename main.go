package main

import (
	"echo/utils"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echo/web"
)

func main() {
	e := echo.New()
	e.Validator = utils.GetCustomValidator()

	e.Use(middleware.Logger())

	// serve pprof endpoints
	e.GET("/debug/pprof", echo.WrapHandler(http.DefaultServeMux))

	e.GET("/", func(c echo.Context) error {
		e.Logger.Debug(fmt.Sprintf("got %s request\n", c.Path()))
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", web.SaveUser)
	e.GET("/users/:id", web.GetUser)
	e.PUT("/users/:id", web.UpdateUser)
	e.DELETE("/users/:id", web.DeleteUser)

	e.Logger.Fatal(e.Start(":8888"))
}
