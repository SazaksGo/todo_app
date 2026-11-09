package main

import (
	"fmt"

	"github.com/SazaksGo/todo_app/http"
	"github.com/SazaksGo/todo_app/task"
)

func main() {
	todoList := task.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Fail to start server ")
	}

}
