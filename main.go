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
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "wrong entry of data"})
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
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggle_todo_completed(context *gin.Context) {
	id := context.Param("id")
	todo, err := get_todo_by_id(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
func delete_todo(context *gin.Context) {
	id := context.Param("id")
	for i, todo := range todos {
		if id == todo.ID {

			copy(todos[i:], todos[i+1:])
			todos = todos[:len(todos)-1]
			context.IndentedJSON(http.StatusOK, gin.H{"message": "todo is succesfully deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
func update_todo_item(context *gin.Context) {
	id := context.Param("id")
	for i := range todos {
		if todos[i].ID == id {
			var item string
			if err := context.BindJSON(&item); err != nil {
				return
			}
			todos[i].Item = item
			context.IndentedJSON(http.StatusOK, gin.H{"message": "item is updated succesfully"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}

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
	router.DELETE("/todos/:id", delete_todo)
	router.PUT("/todos/:id", update_todo_item)
	router.Run("localhost:8080")
}
