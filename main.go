package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"my-imgur/router"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("./.env.development")
	if err != nil {
		fmt.Println("can't load .env.development")
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	_router := router.NewRouter()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/image", _router.UploadPicture)
	e.GET("/obsidian/:path", _router.GetPublicThumbnailLink)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
