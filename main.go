package main	

import (
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/student", handler.CreateStudent)
	router.GET("/student/:id", handler.GetStudent)
	router.PUT("/student/:id", handler.UpdateStudent)
	router.DELETE("/student/:id", handler.DeleteStudent)

	router.Run()
}
