package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all categories
func GetCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"categories": categories,
	})
}

// Create category
func CreateCategory(c *gin.Context) {
	var input models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name: input.Name,
	}

	config.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{
		"message":  "category created",
		"category": category,
	})
}

// Get category by ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// Update category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = input.Name
	config.DB.Save(&category)

	c.JSON(http.StatusOK, category)
}

// Delete category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	config.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
