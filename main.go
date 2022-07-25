package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBFILE = "./database/todos.db"

var DB, error = gorm.Open(sqlite.Open(DBFILE), &gorm.Config{})

type todo struct {
	ID        int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL" jason:"id"`
	Item      string `jason:"item"`
	Completed bool   `jason:"completed"`
}

func get_todos(context *gin.Context) {
	var all_todos []todo
	DB.Find(&all_todos)
	context.IndentedJSON(http.StatusOK, all_todos)
}

func add_todo(context *gin.Context) {
	var new_todo_item string
	if err := context.BindJSON(&new_todo_item); err != nil {
		return
	}
	if len(new_todo_item) == 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "item cannot be empty"})
		return
	}
	var new_todo todo
	new_todo.Item = new_todo_item
	DB.Create(&new_todo)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "todo is added succesfully"})
}

func get_todo(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be number"})
		return
	}
	var res_todo todo
	res := DB.First(&res_todo, id)
	if res.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, res_todo)
}

func delete_todo(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
		return
	}
	var new = todo{ID: id}
	res := DB.Delete(&new)
	if res.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "item is deleted succesfully"})
}

func toggle_todo_completed(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
	}
	var res_todo todo
	DB.First(&res_todo, id)
	if res_todo.Item == "" {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	DB.Model(&todo{}).Where("id = ?", id).Update("completed", !res_todo.Completed)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "todo status is changed"})
}

func update_todo_item(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
	}
	var res_todo todo
	DB.First(&res_todo, id)
	if res_todo.Item == "" {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	var item string
	if err := context.BindJSON(&item); err != nil {
		return
	}
	DB.Model(&todo{}).Where("id = ?", id).Update("item", item)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "item is updated succesfully"})
}

func main() {
	if error != nil {
		panic("couldn't connect to database")
	}
	DB.AutoMigrate(&todo{})
	router := gin.Default()
	router.GET("/todos", get_todos)
	router.GET("/todos/:id", get_todo)
	router.PATCH("/todos/:id", toggle_todo_completed)
	router.POST("/todos", add_todo)
	router.DELETE("/todos/:id", delete_todo)
	router.PUT("/todos/:id", update_todo_item)
	router.Static("/swaggerui/", "swagger_ui")
	router.Run(":8080")
}
