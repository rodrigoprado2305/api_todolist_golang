package main

import (
	"api_todolist_golang/config"
	"api_todolist_golang/controller"
	"api_todolist_golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	config.ConnectDB()

	// Criar um novo roteador Gin
	r := gin.Default()

	// Criar um controlador de tarefas passando a conexão com o banco
	taskController := &controller.TaskController{DB: config.DB}

	// Configurar as rotas
	routes.SetupRoutes(r, taskController)

	// Iniciar o servidor na porta 8080
	r.Run(":8080")
}
