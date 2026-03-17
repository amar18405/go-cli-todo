package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type createTodoRequest struct {
	Title string `json:"title"`
}

type updateTodoRequest struct {
	Title     *string `json:"title"`
	Completed *bool  `json:"completed"`
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from todo Api")
	})

	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoByIDHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func todoByIDHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	// if len(parts) != 3 || parts[2] == "" {
	// 	http.Error(w, "Invalid URL ", http.StatusBadRequest)
	// 	return
	// }

	// id, err := strconv.Atoi(parts[2])
	// if err != nil {
	// 	http.Error(w, "Invalid IDProvided", http.StatusBadRequest)
	// 	return
	// }

	// var req updateTodoRequest
	// err = json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	http.Error(w, "Invalid Request Body", http.StatusBadRequest)
	// 	return
	// }

	// for i, todo := range todos {
	// 	if todo.ID == id {
	// 		if req.Title != "" {
	// 			todos[i].Title = req.Title
	// 		}
	// 		if req.Completed != nil {
	// 			todos[i].Completed = *req.Completed
	// 		}
	// 		saveTodosToFile()

	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(todos[i])
	// 		return
	// 	}
	// }
	// http.Error(w, "Todo Not Found", http.StatusNotFound)

	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])

	if err != nil{
		http.Error(w, " Invalid todo id ", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodPost:
		var req updateTodoRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "couldnt decode the request", http.StatusBadRequest)
		}

		for i, abcd := range todos{
			if abcd.ID == id{
				if req.Title != nil {
					todos[i].Title = *req.Title
				}
				if req.Completed != nil {
					todos[i].Completed = *req.Completed
				}
				saveTodosToFile()
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(todos[i])
				return
			}
		} 	
		http.Error(w,"No todo with given id was found", http.StatusBadRequest)
		
	case http.MethodDelete:
		for i, abcd := range todos{
			if abcd.ID == id{
				todos = append(todos[:i], todos[i+1:]...)
				saveTodosToFile()

				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w,"given Id was not found", http.StatusBadRequest)
	
	default:
		http.Error(w,"wrong method selected", http.StatusMethodNotAllowed)
	}
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
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
			return
		}

		todo := Todo{
			ID:        nextID,
			Title:     req.Title,
			Completed: false,
		}

		nextID++

		todos = append(todos, todo)
		saveTodosToFile()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todos)

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)

	}

}
