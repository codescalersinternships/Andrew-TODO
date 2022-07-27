package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBFILE = "./database/todos.db"

type todo struct {
	ID        int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL" jason:"id"`
	Item      string `jason:"item"`
	Completed bool   `jason:"completed"`
}
type Server struct {
	DB *gorm.DB
}

func (s *Server) get_todos(context *gin.Context) {
	var all_todos []todo
	s.DB.Find(&all_todos)
	context.IndentedJSON(http.StatusOK, all_todos)
}

func (s *Server) add_todo(context *gin.Context) {
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
	s.DB.Create(&new_todo)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "todo is added succesfully"})
}

func (s *Server) get_todo(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be number"})
		return
	}
	var res_todo todo
	res := s.DB.First(&res_todo, id)
	if res.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, res_todo)
}

func (s *Server) delete_todo(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
		return
	}
	var new = todo{ID: id}
	res := s.DB.Delete(&new)
	if res.RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "item is deleted succesfully"})
}

func (s *Server) toggle_todo_completed(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
	}
	var res_todo todo
	res := s.DB.First(&res_todo, id)
	if res.RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	s.DB.Model(&todo{}).Where("id = ?", id).Update("completed", !res_todo.Completed)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "todo status is changed"})
}

func (s *Server) update_todo_item(context *gin.Context) {
	id_s := context.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request id should be num"})
	}
	var res_todo todo
	res := s.DB.First(&res_todo, id)
	if res.RowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	var item string
	if err := context.BindJSON(&item); err != nil {
		return
	}
	s.DB.Model(&todo{}).Where("id = ?", id).Update("item", item)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "item is updated succesfully"})
}

func main() {
	s := Server{}
	var err error
	s.DB, err = gorm.Open(sqlite.Open(DBFILE), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to database")
	}
	s.DB.AutoMigrate(&todo{})
	router := gin.Default()
	router.GET("/todos", s.get_todos)
	router.GET("/todos/:id", s.get_todo)
	router.PATCH("/todos/:id", s.toggle_todo_completed)
	router.POST("/todos", s.add_todo)
	router.DELETE("/todos/:id", s.delete_todo)
	router.PUT("/todos/:id", s.update_todo_item)
	router.Static("/swaggerui/", "swagger_ui")
	router.Run(":8080")
}
