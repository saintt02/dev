package models

import "time"

// Task represent a single task in the task manager
type Task struct {
	ID int
	Name string
	Description string
	Date time.Time
	Completed bool
}