package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `jason:"id"`
	Item      string `jason:"item"`
	Completed bool   `jason:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "do homeword", Completed: false},
	{ID: "2", Item: "exercise", Completed: false},
	{ID: "3", Item: "drink water", Completed: false},
}

func get_todos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func add_todo(context *gin.Context) {
	var new_todo todo
	if err := context.BindJSON(&new_todo); err != nil {
		return
	}
	todos = append(todos, new_todo)
	context.IndentedJSON(http.StatusCreated, todos)
}

func get_todo(context *gin.Context) {
	id := context.Param("id")
	todo, err := get_todo_by_id(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggle_todo_completed(context *gin.Context) {
	id := context.Param("id")
	todo, err := get_todo_by_id(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

//func
func get_todo_by_id(id string) (*todo, error) {
	for i, to_do := range todos {
		if to_do.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("invalid ID")
}

func main() {
	router := gin.Default()
	router.GET("/todos", get_todos)
	router.GET("/todos/:id", get_todo)
	router.PATCH("/todos/:id", toggle_todo_completed)
	router.POST("/todos", add_todo)
	router.Run("localhost:9090")
}
