package main

import (
	"docker-go/controller"
	"docker-go/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//DB接続確認
func connect(c echo.Context) error {
	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", connect)
	e.GET("/todos", controller.GetTodos)
	e.GET("/todo/:id", controller.GetTodo)
	e.POST("/todo", controller.CreateTodo)
	e.PUT("/todo/:id", controller.UpdateTodo)
	e.DELETE("/todo/:id", controller.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))
}