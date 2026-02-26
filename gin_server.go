package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func startGinServer(){
	r := gin.Default()

	r.GET("/todos", getTodosGin)

	r.Run(":8080")
}

func getTodosGin(c *gin.Context){
	c.JSON(http.StatusOK, todos)
}
