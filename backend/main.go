/*
 TODO:write something...
*/

package main

import (
	"net/http"
	"os"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mock-mock/mockmock-meter/backend/command"
	"github.com/mock-mock/mockmock-meter/backend/utils"

	dao "github.com/mock-mock/mockmock-meter/backend/dao"
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
	
	//hostname取得テスト
	//hostname: USER-no-MacBook-Air.local
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("hostname:", name)

	//currentDir取得テスト
	//dir: /Users/user/go/src/github.com/mock-mock/mockmock-meter/backend
	p, _ := os.Getwd()
    fmt.Println("dir:", p)

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

	e.GET("/api/users", func(c echo.Context) error {
		data := dao.GetTest()
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/api/DBusers", func(c echo.Context) error {
		data := dao.GetFromDB()
		return c.JSON(http.StatusOK, data)
	})

	e.POST("/api/mock/start", func(c echo.Context) error {
		// get request param
		requestParam := utils.ParseRequest(c)
		data := command.Start(requestParam)
		return c.JSON(http.StatusOK, data)
	})

	e.POST("/api/mock/end", func(c echo.Context) error {
		// get request param
		requestParam := utils.ParseRequest(c)
		data := command.End(requestParam)
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