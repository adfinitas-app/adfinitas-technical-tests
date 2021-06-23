package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func displayAllRoutes(c echo.Context) error {
	return c.JSON(http.StatusOK, []struct {
		Name   string `json:"name"`
		Path   string `json:"path"`
		Method string `json:"method"`
	}{
		{Name: "displayAllRoutes", Path: "/", Method: "GET"},
	})
}

func computeFibonacci(c echo.Context) error {
	n, _ := strconv.ParseInt(c.QueryParam("n"), 10, 64)

	go func() {
		fmt.Println(fibonacci(n))
	}()

	return c.String(http.StatusOK, strconv.FormatInt(n, 10))
}

func fibonacci(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	e := echo.New()

	e.GET("/displayAllRoutes", displayAllRoutes).Name = "displayAllRoutes"
	e.GET("/computeFibonacci", computeFibonacci)

	e.Logger.Fatal(e.Start(":1323"))
}
