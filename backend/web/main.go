/*
 TODO:write something...
*/

package main

import (
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
	e.Static("/", "/app/frontend")

	// サーバー起動
	//e.Start(":8080")
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	//e.Logger.Fatal(e.Start(":8080"))
}
