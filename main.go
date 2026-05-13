package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func listTodo() {

	if len(todoList) == 0 {
		fmt.Println("No Todo Item Found")
	}

	for _, todo := range todoList {

		status := "Pending"
		if todo.IsCompleted {
			status = "Completed"
		}

		fmt.Printf("%v - %v - %v\n", todo.Id, todo.Title, status)
	}
}

func addTodo() {
	fmt.Println("Enter the title of Todo:")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	todo := Todo{Id: nextID, Title: title, IsCompleted: false}
	todoList = append(todoList, todo)
	nextID++
}

func completeTodo() {
	fmt.Println("Enter Todo ID to mark it completed:")
	var id int
	fmt.Scan(id)

	for _, todo := range todoList {
		if id == todo.Id {
			todo.IsCompleted = true
			fmt.Println("Todo with", id, "is marked as Completed")
			return
		}
	}

	fmt.Println("Todo with", id, "not found")
}

func deleteTodo(){
	fmt.Println("Enter Todo is to delete it:")
	var id int 
	fmt.Scan(id)
	for index,todo := range todoList{
		if id==todo.Id{
			todoList = append(todoList[:index],todoList[index+1:]...)
			fmt.Println("Todo Item with ID",id, "is deleted")
			return
		}
	}
	fmt.Println("Todo with", id ,"not found")
}