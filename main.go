package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MyResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/", func(c echo.Context) error {
		res := new(MyResponse)
		res.Data.Link = "https://xxxxxx.com/xxx.png"
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
