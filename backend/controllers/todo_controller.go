package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ----------------------------
// GET /api/todos (Pagination + Search)
// ----------------------------
func GetTodos(c *gin.Context) {
	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")

	offset := (page - 1) * limit

	var todos []models.Todo
	var total int64

	query := config.DB.Model(&models.Todo{}).Preload("Category")

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	query.Limit(limit).Offset(offset).Find(&todos)

	totalPages := (total + int64(limit) - 1) / int64(limit)

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total":        total,
			"total_pages":  totalPages,
		},
	})
}

// ----------------------------
// POST /api/todos (Create)
// ----------------------------
func CreateTodo(c *gin.Context) {
	var input models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
		CategoryID:  input.CategoryID,
		Priority:    input.Priority,
		DueDate:     input.DueDate,
	}

	config.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "todo created", "todo": todo})
}

// ----------------------------
// GET /api/todos/:id
// ----------------------------
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := config.DB.Preload("Category").First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// ----------------------------
// PUT /api/todos/:id (Update)
// ----------------------------
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = input.Title
	todo.Description = input.Description
	todo.CategoryID = input.CategoryID
	todo.Priority = input.Priority
	todo.DueDate = input.DueDate

	config.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// ----------------------------
// DELETE /api/todos/:id
// ----------------------------
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}

// ----------------------------
// PATCH /api/todos/:id/complete (Toggle complete)
// ----------------------------
func ToggleTodoCompleted(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	config.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{
		"message": "status updated",
		"todo":    todo,
	})
}
