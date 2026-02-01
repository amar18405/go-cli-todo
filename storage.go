package main

import(
	"encoding/json";
	"fmt";
	"os"
)

var todos []Todo
var nextID = 1

func loadTodosFromFile(){
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err){
			todos = []Todo{}
			 	return
		}
		fmt.Println("Error reading tasks file : ", err)
	}
	err = json.Unmarshal(data, &todos);
	if err != nil{
		fmt.Println("Error parsing tasks file : ", err)
		todos = []Todo{}
		return
	}

	maxID := 0

	for _, todo := range todos{
		if todo.ID > maxID{
			maxID = todo.ID
		}
	}
	nextID = maxID + 1
}

func saveTodosToFile(){
	data, err := json.MarshalIndent(todos, "", "  ")

	if err != nil {
		fmt.Println("error encoding todo tasks: ", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)

	if err != nil {
		fmt.Println("Error writing tasks file: ", err)
		return
	}
}