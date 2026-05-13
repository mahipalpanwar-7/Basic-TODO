package main

import "fmt"

var todoList []Todo
var nextID = 1

func main() {
	for {
		fmt.Println("TODO App")
		fmt.Println("1. Show Todo list")
		fmt.Println("2. Add item to Todo")
		fmt.Println("3. Mark Todo as Completed")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			listTodo()
		case 2:
			addTodo()
		case 3:
			completeTodo()
		case 4:
			deleteTodo()
		case 5:
			fmt.Println("Bye")
		default:
			fmt.Println("Invalid Input")
		}

	}
}
