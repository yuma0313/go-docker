package model

type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Detail string `json:"detail"`
}