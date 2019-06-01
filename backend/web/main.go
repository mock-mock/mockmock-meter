/*
 TODO:write something...
*/

package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// サーバー起動
	//e.Start(":8080")
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	//e.Logger.Fatal(e.Start(":8080"))
}
