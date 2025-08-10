package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID         int
	Title        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time 
}


type Todos []Todo 

func (todos *Todos) AddTodo(title string)  *Todos{

	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos= append(*todos, todo)
     return  todos
}


func main() {
  
	todos := Todos{}
    
	todos.AddTodo("vath khyoa")
    
	fmt.Println(todos)
  
}