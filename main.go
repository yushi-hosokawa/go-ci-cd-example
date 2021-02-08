package main 

import (
	"github.com/labstack/echo"
	"go-ci-cd-example/math"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, strconv.Itoa(math.Answer()))
	})
	e.Logger.Fatal(e.Start(":1323"))
}