package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var todoList []Todo
var nextID = 1

func main() {
	for {
		fmt.Println("\nTODO App")
		fmt.Println("1. Show Todo list")
		fmt.Println("2. Add item to Todo")
		fmt.Println("3. Mark Todo as Completed")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid number ")
			continue
		}

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
			return
		default:
			fmt.Println("Invalid Input")
		}

	}
}

func listTodo() {

	if len(todoList) == 0 {
		fmt.Println("No Todo Item Found")
		return
	}

	fmt.Println("\nTodo list")
	for _, todo := range todoList {

		status := "Pending"
		if todo.IsCompleted {
			status = "Completed"
		}

		fmt.Printf("%d - %s - %s\n", todo.Id, todo.Title, status)
	}
}

func addTodo() {
	fmt.Println("Enter the title of Todo:")

	reader := bufio.NewReader(os.Stdin)

	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if title == "" {
		fmt.Println("Todo title cannot be empty")
		return
	}

	todo := Todo{
		Id:          nextID,
		Title:       title,
		IsCompleted: false,
	}

	todoList = append(todoList, todo)
	nextID++

	fmt.Println("Item is added successfully")
}

func completeTodo() {
	fmt.Println("Enter Todo ID to mark it completed:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for index, todo := range todoList {
		if id == todo.Id {
			todoList[index].IsCompleted = true

			fmt.Println("Todo with id", id, "is marked as Completed")
			return
		}
	}

	fmt.Println("Todo with id", id, "not found")
}

func deleteTodo() {
	fmt.Println("Enter Todo is to delete it:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	for index, todo := range todoList {

		if id == todo.Id {
			todoList = append(
				todoList[:index],
				todoList[index+1:]...,
			)
			fmt.Println("Todo Item with ID", id, "is deleted successfully")
			return
		}
	}
	fmt.Println("Todo with id ", id, "not found")
}
