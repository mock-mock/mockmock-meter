/*
 TODO:write something...
*/

package main

import (
	"net/http"
	//"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mock-mock/mockmock-meter/backend/command"
	"github.com/mock-mock/mockmock-meter/backend/domain"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	//e.Static("/", "../../frontend/mockmockproto/dist")
	//e.Static("/", "../../frontend/vuetify-material-dashboard-master/dist")
	//e.Static("/", "/app/frontend")
	//e.Static("/", "../../frontend/vuetify-material-dashboard-master/dist")
	e.Static("/", "/app/frontend/vuetify-material-dashboard-master/dist")

	e.GET("/dashboard", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://mockmock-meter-proto.herokuapp.com")
	})

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello!")
	})

	e.GET("/api/healthCheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK!")
	})

	e.POST("/api/mock/start", func(c echo.Context) error {
		// get request param
		req := new(domain.SlackRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		data := command.Start(*req)
		return c.JSON(http.StatusOK, data)
	})

	e.POST("/api/mock/end", func(c echo.Context) error {
		// get request param
		req := new(domain.SlackRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		data := command.End(*req)
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/api/mock/${term}", func(c echo.Context) error {
		data := command.Get()
		return c.JSON(http.StatusOK, data)
	})

	// サーバー起動
	//e.Start(":8080")
	//e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	e.Logger.Fatal(e.Start(":8080"))
}
