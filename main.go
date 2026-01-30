package main

import (
	"fmt"
	"os"
	"strconv"
)

func printUsage(){
	fmt.Println("Usage :")
	fmt.Println("	todo add \"todo description\"")
	fmt.Println("	todo list")
	fmt.Println("	todo done <id>")
	fmt.Println("	todo delete <id>")
}

func main(){
	args := os.Args

	if len(args) < 2 {
		printUsage()
		return
	}

	command := args[1]
	
	switch command{
	case "add":
		if(len(args) < 3){
			fmt.Println("Please provide a task description.")
			return
		}
		addTodo(args[2])

	case "list":
		listTodo()

	case "done":
		if len(args) < 3{
			fmt.Println("Enter a task ID ")
			return
		}

		id, err := strconv.Atoi(args[2])

		if err != nil{
			fmt.Println("Invalid Task ID!")
			return
		}

		markTodoDone(id)

	case "delete":
		if len(args) < 3{
			fmt.Println("Enter a task ID")
			return
		}

		id, err := strconv.Atoi(args[2])

		if err != nil {
			fmt.Println("Invalid Task ID")
			return
		}

		deleteTodo(id)

	default:
		fmt.Println("Unknown command selected : ", command)
		printUsage()
	}
}