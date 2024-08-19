package handlers

import (
	"fmt"
	"task-manager/models"
	"time"
)

// Uso de Slices: Se utiliza un slice para almacenar las tareas, ya que es flexible y fácil de manipular.
var tasks []models.Task

var nextID int

// AddTasks adds a new task to the list
func AddTask(name, description string, date time.Time) {
	task := models.Task {
		ID: nextID,
		Name: name,
		Description: description,
		Date: date,
		Completed: false,
	}
	tasks = append(tasks, task)
	// ID Incremental: El ID se maneja manualmente para evitar duplicados y asegurar la unicidad de cada tarea.
	nextID ++
	// Mensajes Informativos: Cada acción devuelve un mensaje que indica el éxito o el fallo de la operación
	fmt.Println("Task added successfully!")
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d\nName:%s\nDescription: %s\nDate: %s\nStatus: %s\n\n",
		task.ID, task.Name, task.Description, task.Date.Format("2006-01-02"), status)
	}
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found.")
}

func MarkTaskAsCompleted(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed")
			return
		}
	}
	fmt.Println("Task not found.")
}
