package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PrintTeste(texto string) {
	fmt.Println(texto)
}

func Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
