package main

import "fmt"

func addTodo(title string){
	todo := Todo{
		ID : nextID,
		Title : title,
		Completed: false,
	}

	nextID++

	todos = append(todos,todo)

	fmt.Println("Added task: ", todo.Title)
}

func listTodo(){
	if len(todos) == 0{
		fmt.Println("No tasks found!!!")
		return
	}

	for _, i := range todos {
		status := " "
		if(i.Completed){
			status = " Done "
		}
		
		fmt.Printf("[%s] %d: %s \n", status, i.ID, i.Title)
	}
	return
}