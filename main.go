package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	//baixar o gim pelo powershell
	//go get github.com/gin-gonic/gin

	_ "github.com/go-sql-driver/mysql" //underscore, Importação anônima com underscore
	//O underscore (_) é necessário aqui porque o pacote mysql está sendo importado para registrar o driver, mas não é usado diretamente no código.
	//go get github.com/go-sql-driver/mysql
	//O uso do _ antes da importação do driver MySQL (_ "github.com/go-sql-driver/mysql") está correto, pois o driver é registrado implicitamente
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:Banco12345*@tcp(localhost:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Criar tabela se não existir
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	//POST para criar uma nova tarefa:
	//curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title": "Minha nova tarefa", "description": "Descrição da tarefa", "completed": false}'

	//GET para obter todas as tarefas:
	//curl http://localhost:8080/tasks

	//GET para obter uma tarefa específica (substitua {id} pelo ID real):
	//curl http://localhost:8080/tasks/1

	//DELETE para excluir uma tarefa (substitua {id} pelo ID real):
	//curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title": "Tarefa atualizada", "description": "Nova descrição", "completed": true}'

	r.Run(":8080")
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	task.ID = int(id)
	c.JSON(201, task)
}

func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(200, tasks)
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	err := db.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE tasks SET title=?, description=?, completed=? WHERE id=?", task.Title, task.Description, task.Completed, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, task)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted"})
}
