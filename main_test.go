package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func MakeTempFile(t testing.TB) string {
	f, err := os.CreateTemp("", "go-sqlite-test")
	defer f.Close()
	if err != nil {
		t.Fatalf("Error making temp file: %q", err.Error())
	}
	return f.Name()
}
func TestAddTodo(t *testing.T) {
	file := MakeTempFile(t)
	defer os.Remove(file)
	var app App
	app.model.GetConnection(file)
	app.model.addTodo("first todo")
	app.model.addTodo("second todo")

	t.Run("testing add todo created status ", func(t *testing.T) {
		router := SetUpRouter()
		router.POST("/todos", app.addTodoHandler)
		new_item := "this is new task"
		json_item, _ := json.Marshal(new_item)
		req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(json_item))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusCreated {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})

	t.Run("testing add todo bad request empty item  ", func(t *testing.T) {
		router := SetUpRouter()
		router.POST("/todos", app.addTodoHandler)
		new_item := ""
		json_item, _ := json.Marshal(new_item)
		req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(json_item))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

}

func TestGetTodo(t *testing.T) {
	file := MakeTempFile(t)
	defer os.Remove(file)
	var app App
	app.model.GetConnection(file)
	app.model.addTodo("first todo")
	app.model.addTodo("second todo")
	t.Run("testing get todo  (OK) with valid id", func(t *testing.T) {
		router := SetUpRouter()
		router.GET("/todos/:id", app.getTodoHandler)
		req, _ := http.NewRequest(http.MethodGet, "/todos/2", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusAccepted {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusAccepted)
		}
	})

	t.Run("testing get todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		router.GET("/todos/:id", app.getTodoHandler)
		req, _ := http.NewRequest(http.MethodGet, "/todos/invalid", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing get todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
		router.GET("/todos/:id", app.getTodoHandler)
		req, _ := http.NewRequest(http.MethodGet, "/todos/6", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestDeleteTodo(t *testing.T) {
	file := MakeTempFile(t)
	defer os.Remove(file)
	var app App
	app.model.GetConnection(file)
	app.model.addTodo("first todo")
	app.model.addTodo("second todo")
	t.Run("testing delete todo  (Ok) with valid id", func(t *testing.T) {
		router := SetUpRouter()
		router.DELETE("/todos/:id", app.deleteTodoHandle)
		req, _ := http.NewRequest(http.MethodDelete, "/todos/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNoContent {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNoContent)
		}
	})
	t.Run("testing delete todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		router.DELETE("/todos/:id", app.deleteTodoHandle)
		req, _ := http.NewRequest(http.MethodDelete, "/todos/invalid", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing delete todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
		router.DELETE("/todos/:id", app.deleteTodoHandle)
		req, _ := http.NewRequest(http.MethodDelete, "/todos/10", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestUpdateTodo(t *testing.T) {
	file := MakeTempFile(t)
	defer os.Remove(file)
	var app App
	app.model.GetConnection(file)
	app.model.addTodo("first todo")
	app.model.addTodo("second todo")
	t.Run("testing update todo  (ok) with valid", func(t *testing.T) {
		router := SetUpRouter()
		router.PUT("/todos/:id", app.updateTodoHandler)
		new_item := "this is updated task"
		json_item, _ := json.Marshal(new_item)
		req, _ := http.NewRequest(http.MethodPut, "/todos/1", bytes.NewBuffer(json_item))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusCreated {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})
	t.Run("testing update todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		router.PUT("/todos/:id", app.updateTodoHandler)
		new_item := "this is updated task"
		json_item, _ := json.Marshal(new_item)
		req, _ := http.NewRequest(http.MethodPut, "/todos/invalidd", bytes.NewBuffer(json_item))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing update todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
		router.PUT("/todos/:id", app.updateTodoHandler)
		new_item := "this is updated task"
		json_item, _ := json.Marshal(new_item)
		req, _ := http.NewRequest(http.MethodPut, "/todos/10", bytes.NewBuffer(json_item))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}
