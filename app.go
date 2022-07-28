package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL" jason:"id"`
	Item      string `jason:"item"`
	Completed bool   `jason:"completed"`
}

type Model struct {
	db *gorm.DB
}

func (m *Model) getConnection(file string) {
	var err error
	m.db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		errstr := fmt.Sprintf("couldnt connect to database in this file: %s\n ERROR: %s", file, err)
		panic(errstr)
	}
	m.db.AutoMigrate(&Todo{})
}

func (m *Model) getTodo(id int) (Todo, error) {
	var res_todo Todo
	res := m.db.First(&res_todo, id)
	return res_todo, res.Error
}

func (m *Model) deleteTodo(id int) error {
	var new = Todo{ID: id}
	res := m.db.Delete(&new)
	if res.RowsAffected == 0 {
		return errors.New("ID is not found")
	}
	return nil
}

func (m *Model) addTodo(new_todo_string string) error {
	var new_todo Todo
	new_todo.Item = new_todo_string
	res := m.db.Create(&new_todo)
	return res.Error
}

func (m *Model) getTodos() ([]Todo, error) {
	var all_todos []Todo
	res := m.db.Find(&all_todos)
	return all_todos, res.Error
}

func (m *Model) updateTodo(id int, todo_item string) error {
	res := m.db.Model(&Todo{}).Where("id = ?", id).Update("item", todo_item)
	return res.Error
}
