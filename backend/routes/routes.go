package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	// ---------------------
	// Category Routes
	// ---------------------
	api.GET("/categories", controllers.GetCategories)
	api.POST("/categories", controllers.CreateCategory)
	api.GET("/categories/:id", controllers.GetCategoryByID)
	api.PUT("/categories/:id", controllers.UpdateCategory)
	api.DELETE("/categories/:id", controllers.DeleteCategory)

	// ---------------------
	// Todo Routes
	// ---------------------
	api.GET("/todos", controllers.GetTodos)
	api.POST("/todos", controllers.CreateTodo)
	api.GET("/todos/:id", controllers.GetTodoByID)
	api.PUT("/todos/:id", controllers.UpdateTodo)
	api.DELETE("/todos/:id", controllers.DeleteTodo)
	api.PATCH("/todos/:id/complete", controllers.ToggleTodoCompleted)
}
