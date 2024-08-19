package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task-manager/handlers"
	"time"
)

func main() {
	for {
		fmt.Println("Task Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Delete Tasks")
		fmt.Println("4. Mark Task as Completed")
		fmt.Println("5. Exit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddTaskMenu()
		case 2: 
			handlers.ListTasks()
		case 3:
			DeleteTaskMenu()
		case 4:
			MarkTaskAsCompletedMenu()
		case 5: 
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}

	}
}

func AddTaskMenu() {
	reader := bufio.NewReader(os.Stdin)

	// reader.ReadString('\n'): Esta función lee toda la línea de entrada hasta que encuentra un salto de línea (\n). Esto significa que la descripción podrá incluir espacios.

	// strings.TrimSpace(): Este método se usa para eliminar cualquier espacio en blanco o saltos de línea al principio o al final de la cadena. Es útil para limpiar la entrada antes de procesarla.
	fmt.Print("Enter task name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Enter task due date (YYYY-MM-DD): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)

	// ("2006-01-02"): Es un layout específico utilizado por Go como referencia para saber cómo interpretar una cadena de texto como fecha. Debes usar este formato exacto cuando llames a time.Parse() para que el analizador sepa qué partes de la cadena corresponden a los días, meses y años.
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
			fmt.Println("Invalid date format, please use YYYY-MM-DD.")
			return
	}
	
	handlers.AddTask(name, description, date)
}

func DeleteTaskMenu() {
	var id int 
	fmt.Print("Enter task ID to delete: ")
	fmt.Scanln(&id)
	handlers.DeleteTask(id)
}

func MarkTaskAsCompletedMenu() {
	var id int
	fmt.Print("Enter task ID to mark as completed: ")
	fmt.Scanln(&id)
	handlers.MarkTaskAsCompleted(id)
}