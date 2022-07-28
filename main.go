package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const DBFILE = "./database/todos.db"

type App struct {
	model Model
}

func (app *App) getTodosHandler(context *gin.Context) {
	var all_todos []Todo
	all_todos, err := app.model.getTodos()
	if err == nil {
		context.IndentedJSON(http.StatusOK, all_todos)
		return
	}
	//must be error here

}

func (app *App) addTodoHandler(context *gin.Context) {
	var new_todo_item string
	if err := context.BindJSON(&new_todo_item); err != nil {
		return
	}
	if len(new_todo_item) == 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "item cannot be empty"})
		return
	}
	err := app.model.addTodo(new_todo_item)
	if err == nil {
		context.IndentedJSON(http.StatusOK, gin.H{"message": "todo is added succesfully"})
	}
	// there should be error here
}

func (app *App) getTodoHandler(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be number"})
		return
	}

	res_todo, err := app.model.getTodo(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, res_todo)
}

func (app *App) deleteTodoHandle(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
		return
	}
	err2 := app.model.deleteTodo(id)
	if err2 != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "item is deleted succesfully"})
}

func (app *App) updateTodoHandler(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
	}
	updated_todo, err2 := app.model.getTodo(id)
	if err2 != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	var item string
	if err := context.BindJSON(&item); err != nil {
		return
	}
	err3 := app.model.updateTodo(updated_todo.ID, item)
	if err3 == nil {
		context.IndentedJSON(http.StatusOK, gin.H{"message": "item is updated succesfully"})
	}
	//there should be error
}

func main() {
	app := App{}
	app.model.GetConnection(DBFILE)
	router := gin.Default()
	router.GET("/todos", app.getTodosHandler)
	router.GET("/todos/:id", app.getTodoHandler)
	router.POST("/todos", app.addTodoHandler)
	router.DELETE("/todos/:id", app.deleteTodoHandle)
	router.PUT("/todos/:id", app.updateTodoHandler)
	router.Static("/swaggerui/", "swagger_ui")
	router.Run(":8080")
}
