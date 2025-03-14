package controller

import (
	"api_todolist_golang/model"
	"api_todolist_golang/repository"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"log"
	"sync"
)

type TaskController struct {
	DB *sql.DB
}

func logRequest(c *gin.Context) { // Chama a goroutine para log
	log.Printf("Requisição recebida: %s %s", c.Request.Method, c.Request.URL.Path)
}

func logRequestRoutine(c *gin.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Go routine - Requisição recebida: %s %s", c.Request.Method, c.Request.URL.Path)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	logRequest(c)

	var task model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := repository.CreateTask(tc.DB, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.ID = int(id)
	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go logRequestRoutine(c, &wg)

	tasks, err := repository.GetTasks(tc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	logRequest(c)

	id, _ := strconv.Atoi(c.Param("id"))
	task, err := repository.GetTask(tc.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	logRequest(c)

	id, _ := strconv.Atoi(c.Param("id"))
	var task model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = id

	if err := repository.UpdateTask(tc.DB, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	logRequest(c)

	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteTask(tc.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
