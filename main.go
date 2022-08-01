package main

import (
	middlewaree "TO_DO_PROJECT/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const DBFILE = "./database/todos.db"

type App struct {
	model Model
}

func (app *App) CreateApp(file string) {
	app.model.getConnection(file)
}

func (app *App) getTodosHandler(context *gin.Context) {
	var all_todos []Todo
	all_todos, err := app.model.getTodos()
	if err == nil {
		context.IndentedJSON(http.StatusOK, all_todos)
		return
	}
	context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})

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
	todo, err := app.model.addTodo(new_todo_item)
	if err == nil {
		context.IndentedJSON(http.StatusCreated, todo)
		return
	}
	context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
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
	context.IndentedJSON(http.StatusNoContent, gin.H{"message": "item is deleted succesfully"})
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
		context.IndentedJSON(http.StatusAccepted, gin.H{"message": "item is updated succesfully"})
		return
	}
	context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
}

func (app *App) updateTodoStatusHandler(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
	}
	togeled_todo, err2 := app.model.getTodo(id)
	if err2 != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	err3 := app.model.toggleTodo(id, togeled_todo)
	if err3 == nil {
		context.IndentedJSON(http.StatusAccepted, gin.H{"message": "item is toggeled succesfully"})
		return
	}
	context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
}

func main() {
	app := App{}
	app.CreateApp(DBFILE)
	router := gin.New()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	router.Use(cors.New(cors.Config{
		AllowHeaders: []string{"Content-Type", "Access-Control-Allow-Origin"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "DELETE", "GET", "POST", "PATCH"},
	}))

	router.Use(middlewaree.MiddleWare())
	router.GET("/todos", app.getTodosHandler)
	router.GET("/todos/:id", app.getTodoHandler)
	router.POST("/todos", app.addTodoHandler)
	router.PATCH("/todos/:id", app.updateTodoStatusHandler)
	router.DELETE("/todos/:id", app.deleteTodoHandle)
	router.PUT("/todos/:id", app.updateTodoHandler)
	router.Static("/swaggerui/", "swagger_ui")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
