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
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	_router := router.NewRouter()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	g := e.Group("/v1/image")
	g.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:    nil,
		KeyLookup:  "header:Authorization",
		AuthScheme: "Client-ID",
		Validator: func(auth string, c echo.Context) (bool, error) {
			return auth == os.Getenv("CLIENT_ID"), nil
		},
		ErrorHandler:           nil,
		ContinueOnIgnoredError: false,
	}))
	g.POST("", _router.UploadPicture)
	e.GET("/v1/temp-link/obsidian/:fileName", _router.GetPublicThumbnailLink)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
