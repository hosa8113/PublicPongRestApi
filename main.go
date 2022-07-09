package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const UrlBasePath = ""
const KlfDateTimeFormat = "2006-01-02-15:04:05.999999999"

func main() {
	e := echo.New()

	// ping
	e.GET(UrlBasePath+"/ping", func(context echo.Context) error {
		println(time.Now().Format(KlfDateTimeFormat) + " - Got request from " + context.RealIP())
		return context.JSON(http.StatusOK, "pong")
	})

	// hello
	e.GET(UrlBasePath+"/hello/:name", func(context echo.Context) error {
		println(time.Now().Format(KlfDateTimeFormat) + " - Got request from " + context.RealIP())
		name := context.Param("name")
		return context.JSON(http.StatusOK, "Hello "+name)
	})
	e.Logger.Fatal(e.Start(":8200"))
}
