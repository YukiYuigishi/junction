package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var redirect = make(map[string]string)

func Redirect(c echo.Context) error {
	idParam := c.QueryParam("id")
	fmt.Println("test", "id="+idParam)
	if len(idParam) > 32 {
		return c.String(http.StatusNotFound, "id too long")
	}
	redirectUrl, exist := redirect[idParam]
	if !exist {
		return c.String(http.StatusNotFound, "Query Error")
	}
	fmt.Println("redirect url:", redirectUrl)
	return c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

func main() {
	redirect["0"] = "https://github.com/YukiYuigishi/MEISHI/blob/master/PCB/MEISHI2022/README.md"
	redirect["1"] = "https://github.com/YukiYuigishi/MEISHI/"
	fmt.Println("start junction.")
	app := echo.New()
	app.GET("/", Redirect)

	app.Logger.Fatal(app.Start(":1234"))
}
