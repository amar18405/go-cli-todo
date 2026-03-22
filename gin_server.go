package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startGinServer() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(404, "Hello from gin server")
	})
	r.GET("/todos", getTodosGin)
	r.POST("/todos/:id", updateTodoGin)
	r.POST("/todos", createTodosGin)
	r.DELETE("/todos/:id", deleteTodoGin)
	r.Run(":8080")
}

func deleteTodoGin(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "ID couldnt be decoded",
		})
		return
	}

	results, err := db.Exec("DELETE FROM todos WHERE ID = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error" : "Failed to delete TODO",
		})
		return
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error" : "Failed to verify the deletion",
		})
	}
	if rowsAffected < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error" : "NO Row was found for that id",
		})
	}

	c.Status(http.StatusNoContent)
}

func createTodosGin(c *gin.Context) {

	var req createTodoRequest
	err := c.BindJSON(&req)
	if err != nil || req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Request Body",
		})
	}
	var todo Todo

	err = db.QueryRow("INSERT INTO todos(Title, Completed) VALUES ($1, $2) returning id, title, completed", req.Title, false).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to add the task",
		})
		return
	}

	c.JSON(200, todo)
}

func updateTodoGin(c *gin.Context) {
	//1. get id
	idStr := c.Param("id") // it takes id from the context as a parameter. the id comes from the path we specified in line 18

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(400, "ID not found")
		return
	}

	var req updateTodoRequest
	err = c.BindJSON(&req) // BindJSON converts the json type of context to golang struct
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	//check if atleast one field is provided
	if req.Title == nil && req.Completed == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Nothing to update",
		})
		return
	}

	query := "UPDATE todos SET "
	args := []interface{}{}
	argPos := 1

	if req.Title != nil {
		query += fmt.Sprintf("title = $%d", argPos)
		args = append(args, *req.Title)
		argPos++
	}

	if req.Completed != nil {
		if len(args) > 1 {
			query += ", "
		}
		query += fmt.Sprintf("completed = $%d", argPos)
		args = append(args, *req.Completed)
		argPos++
	}
	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, title, completed", argPos)
	args = append(args, id)
	var todo Todo
	err = db.QueryRow(query, args...).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func getTodosGin(c *gin.Context) {
	rows, err := db.Query("Select Title, ID, Completed FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Todos couldn't be fetched",
		})
		return
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Title, &todo.ID, &todo.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Todos couldn't be decoded",
			})
			return
		}
		todos = append(todos, todo)
	}
	c.JSON(200, todos)
}
