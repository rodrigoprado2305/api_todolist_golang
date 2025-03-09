package repository

import (
	"api_todolist_golang/model"
	"database/sql"
)

// Criar uma nova tarefa no banco
func CreateTask(db *sql.DB, task *model.Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Buscar todas as tarefas
func GetTasks(db *sql.DB) ([]model.Task, error) {
	rows, err := db.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Buscar uma única tarefa pelo ID
func GetTask(db *sql.DB, id int) (*model.Task, error) {
	var task model.Task
	err := db.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Atualizar uma tarefa
func UpdateTask(db *sql.DB, task *model.Task) error {
	_, err := db.Exec("UPDATE tasks SET title=?, description=?, completed=? WHERE id=?", task.Title, task.Description, task.Completed, task.ID)
	return err
}

// Deletar uma tarefa
func DeleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}
