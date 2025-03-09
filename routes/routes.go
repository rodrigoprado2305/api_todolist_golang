package routes

import (
	"api_todolist_golang/controller"

	"github.com/gin-gonic/gin"
	//go get github.com/gin-gonic/gin
)

func SetupRoutes(router *gin.Engine, taskController *controller.TaskController) {
	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks", taskController.GetTasks)
	router.GET("/tasks/:id", taskController.GetTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)
}
