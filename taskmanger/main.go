package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	ID         int
	Title        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time 
}


// type Todos []Todo 

// func (todos *Todos) AddTodo(title string)  *Todos{

// 	todo := Todo{
// 		Title: title,
// 		Completed: false,
// 		CreatedAt: time.Now(),
// 		CompletedAt: nil,
// 	}

// 	*todos= append(*todos, todo)
//      return  todos
// }

type Todos struct{
	Tasks []Todo `json:"tasks"`
}

func (todos *Todos) AddTodo(title string) *Todos {
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	todos.Tasks= append(todos.Tasks, todo)
	return todos
}


func (todos *Todos) SaveToFile(filename string) error {
     data, err := json.MarshalIndent(todos, "", " ")

	 if err !=nil {
		return  err
	 }

	 return  os.WriteFile(filename,data,0644)
}

func (todos *Todos) LoadFromFile(filename string) error {
	file,err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return  json.Unmarshal(file,todos)
}

func main() {
    add := flag.String("add", "", "Add a new task")
	flag.Parse()

	todos := Todos{}
    err:= todos.LoadFromFile("test1")
	   if err != nil {
        fmt.Println("Error loading tasks:", err)
        os.Exit(1)
    }

    switch {
      case *add != "":
        todos.AddTodo(*add)
      default:
        fmt.Println("No command provided. Use -h for help.")
     } 

	
	err = todos.SaveToFile("test1")
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        os.Exit(1)
    }
 
}

