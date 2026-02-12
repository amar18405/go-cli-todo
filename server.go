package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type createTodoRequest struct{
	Title string `json:"title"`
}

func startServer(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintln(w,"Hello from todo Api")
	})

	http.HandleFunc("/todos", todosHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request){
	// //only allow get methods
		
	// if r.Method != http.MethodGet {
	// 	http.Error(w,"Method Not Allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")

	// err := json.NewEncoder(w).Encode(todos)

	// if err != nil {
	// 	http.Error(w,"failed to encode todos", http.StatusInternalServerError)
	// 	return
	// }

	switch r.Method {
	case http.MethodGet:
		//set http response header to indicate that response will be in json format
		w.Header().Set("Content-Type", "application/json")
		//converts todos to json and writes it in the response
		json.NewEncoder(w).Encode(todos)
		
	case http.MethodPost:
		var req createTodoRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return;
		}

		todo := Todo{
			ID : nextID,
			Title: req.Title,
			Completed: false,
		}

		nextID++

		todos = append(todos, todo)
		saveTodosToFile()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todos)

	default:
		http.Error(w,"invalid method", http.StatusMethodNotAllowed)

	}

}