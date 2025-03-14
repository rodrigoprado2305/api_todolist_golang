package routes

import (
	"api_todolist_golang/controller"
	"api_todolist_golang/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, taskController *controller.TaskController) {
	router.POST("/login", controller.GenerateToken)

	// Rotas protegidas
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/tasks", taskController.CreateTask)
		protected.GET("/tasks", taskController.GetTasks)
		protected.GET("/tasks/:id", taskController.GetTask)
		protected.PUT("/tasks/:id", taskController.UpdateTask)
		protected.DELETE("/tasks/:id", taskController.DeleteTask)
	}
}
