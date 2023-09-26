package main

import (
	"docker-go/controller"
	"docker-go/model"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Detail string `json:"detail"`
}

type allTodos []todo

var todos = allTodos{
	{
		ID:     "1",
		Title:  "買い物",
		Detail: "食料品を買う",
	},
	{
		ID:     "2",
		Title:  "プロジェクト",
		Detail: "新しいプロジェクトの計画を立てる",
	},
	{
		ID:     "3",
		Title:  "ジョギング",
		Detail: "公園でジョギングする",
	},
}

func getTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func getTodo(c echo.Context) error {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.ID == id {
			return c.JSON(http.StatusOK, todo)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "TODOが見つかりません")
}

func postTodo(c echo.Context) error {
	var newTodo todo
	json.NewDecoder(c.Request().Body).Decode(&newTodo)
	todos = append(todos, newTodo)

	return c.JSON(http.StatusOK, todos)
}

func updateTodo(c echo.Context) error {
	urlId := c.Param("id")

	var updateTodo todo
	c.Bind(&updateTodo)

	for i, todo := range todos {
		todoId := todo.ID

		if todoId == urlId {
			todos[i].Title = updateTodo.Title
			todos[i].Detail = updateTodo.Detail
			return c.JSON(http.StatusOK, todos)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "TODOが見つかりません")
}

func deleteTodo(c echo.Context) error {
	id := c.Param("id")

	for i, todo := range todos {

		if todo.ID == id {
			todos = append(todos[:i], todos[i + 1:]...)
			return c.JSON(http.StatusOK, todos)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "TODOが見つかりません")
}

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

	e.GET("/", connect)
	e.GET("/todos", controller.GetTodos)
	e.GET("/todo/:id", controller.GetTodo)
	e.POST("/todo", controller.CreateTodo)
	e.PUT("/todo/:id", controller.UpdateTodo)
	e.DELETE("/todo/:id", controller.DeleteTodo)
	e.Logger.Fatal(e.Start(":8080"))
}