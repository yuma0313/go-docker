package controller

import (
	"docker-go/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	todos := []model.Todo{}
	if err := c.Bind(&todos); err != nil {
		return err
	}

	model.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	id := c.Param("id")
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	model.DB.Take(&todo, id)
	return c.JSON(http.StatusOK, todo)
}

func CreateTodo(c echo.Context) error {
	todo := model.Todo{}
	//BindメソッドでリクエストボディからJSONデータを取得する
	if err := c.Bind(&todo); err != nil {
		return err
	}

	model.DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	model.DB.Save(&todo)
	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	todos := []model.Todo{}
	if err := c.Bind(&todos); err != nil {
		return err
	}

	model.DB.Delete(&todos, id)
	model.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}