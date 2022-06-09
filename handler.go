package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateStudent(c *gin.Context) {
	var student Student

	if err := c.BindJSON(&student); err != nil {
		fmt.Printf("failed to bind student: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	h.storage.Insert(&student)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": student.ID,
	})
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	var student Student

	if err := c.BindJSON(&student); err != nil {
		fmt.Printf("failed to bind student: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	h.storage.Update(id, student)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": student.ID,
	})
}

func (h *Handler) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	student, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get student %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *Handler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		
		return
	}

	h.storage.Delete(id)

	c.String(http.StatusOK, "student deleted")
}