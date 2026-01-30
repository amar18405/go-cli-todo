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

func markTodoDone(id int){
	for i, todo := range todos{
		if todo.ID == id {
			if todo.Completed {
				fmt.Println("Task already completed : ", todo.Title)
				return
			}

			todos[i].Completed = true
			fmt.Println("Task marked as done : ", todos[i].Title)
			return
		}
	}
	fmt.Println("No task found with ID : ", id)
	return
}

func deleteTodo(id int){
	for i, todo := range todos{
		if todo.ID == id{
			todos = append(todos[:i],todos[i+1:]... )
			fmt.Println("Task deleted : ", id)
			return
		}
	}
	fmt.Println("No task found with id :", id)
	return
}