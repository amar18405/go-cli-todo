package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func startGinServer(){
	r := gin.Default()
	r.GET("/", func (c * gin.Context)  {
		c.String(404, "Hello from gin server")
	})
	r.GET("/todos", getTodosGin)
	r.POST("/todos/:id", updateTodoGin)
	r.POST("/todos", createTodosGin)
	r.Run(":8080")
}

func createTodosGin(c *gin.Context){

	var req createTodoRequest
	err := c.BindJSON(&req)
	if err != nil || req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : "Invalid Request Body",
		})
	}
	todo := Todo{
		ID : nextID,
		Title: req.Title,
		Completed: false,
	}

	todos = append(todos, todo)
	nextID++
	saveTodosToFile()

	c.JSON(200, todo)
}

func updateTodoGin(c *gin.Context){
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
			"error" : "Invalid Request Body",
		})
		return
	}

	for i, todo := range todos{
		if todo.ID == id{
			if req.Title != nil{
				todos[i].Title = *req.Title
			}
			if req.Completed != nil{
				todos[i].Completed = *req.Completed
			}
			saveTodosToFile()

			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error" : "Todo not found",
	})
}

func getTodosGin(c *gin.Context){
	c.JSON(http.StatusOK, todos)
}
