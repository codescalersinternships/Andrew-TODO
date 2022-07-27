package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	var s Server
	s.DB, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	s.DB.AutoMigrate(&todo{})
	s.DB.Create(&todo{ID: 1, Item: "first todo", Completed: false})
	s.DB.Create(&todo{ID: 2, Item: "second todo", Completed: false})

	t.Run("testing add todo OK status ", func(t *testing.T) {
		router := SetUpRouter()
		router.POST("/todos", s.add_todo)
		new_item := "this is new task"
		json_item, _ := json.Marshal(new_item)
		req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(json_item))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusOK {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("testing add todo bad request empty item  ", func(t *testing.T) {
		router := SetUpRouter()
		router.POST("/todos", s.add_todo)
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
	var s Server
	s.DB, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	s.DB.AutoMigrate(&todo{})
	s.DB.Create(&todo{ID: 1, Item: "first todo", Completed: false})
	s.DB.Create(&todo{ID: 2, Item: "second todo", Completed: false})

	t.Run("testing get todo  (OK) with valid id", func(t *testing.T) {
		router := SetUpRouter()
		router.GET("/todos/:id", s.get_todo)
		req, _ := http.NewRequest(http.MethodGet, "/todos/2", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusOK {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("testing get todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodGet, "/todos/invalid", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing get todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
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
	t.Run("testing delete todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodDelete, "/todos/invalid", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing delete todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodDelete, "/todos/-10", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestUpdateTodo(t *testing.T) {
	t.Run("testing update todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodPut, "/todos/invalidd", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing update todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodPut, "/todos/-10", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestToggleTodo(t *testing.T) {
	t.Run("testing toggle todo  (bad request) with invalid id", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodPatch, "/todos/invalid", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
	t.Run("testing toggle todo  (status not found) with non existent id ", func(t *testing.T) {
		router := SetUpRouter()
		req, _ := http.NewRequest(http.MethodPatch, "/todos/-10", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		status := response.Code
		if status != http.StatusNotFound {
			t.Errorf("Returned Wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}
