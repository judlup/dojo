package main

import (
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Id    int
	Title string
	Done  bool
}

func logic(cmd string, value string, tasks []Task) []Task {
	id := len(tasks) + 1
	switch cmd {
	case "add":
		id++
		task := Task{
			Id:    id,
			Title: value,
			Done:  false,
		}
		if value == "" {
			log("Task title cannot be empty")
		}
		tasks = append(tasks, task)
		log(fmt.Sprintf("Created Task: %+v", task))
	case "list":
		log("Todo List Tasks")
		for _, task := range tasks {
			log(fmt.Sprintf("Task ID: %d, Title: %s, Done: %v", task.Id, task.Title, task.Done))
		}
	case "delete":
		taskId, err := strconv.Atoi(value)
		if err != nil {
			log("Invalid Task ID")
		}
		for i, task := range tasks {
			if task.Id == taskId {
				tasks = append(tasks[:i], tasks[i+1:]...)
				log(fmt.Sprintf("Deleted Task ID: %d", taskId))
				break
			}
		}
		log(fmt.Sprintf("Remaining Tasks: %+v", tasks))
	default:
		log(fmt.Sprintf("Unknown command: %s", cmd))
	}
	return tasks
}

func log(message string) {
	fmt.Println(message)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log("No tasks action provided")
		return
	}
	cmd := args[0]
	tasks := []Task{
		{Id: 1, Title: "Sample Task 1", Done: false},
		{Id: 2, Title: "Sample Task 2", Done: true},
	}
	value := ""
	if len(args) >= 2 {
		value = args[1]
	} else {
		value = ""
	}
	tasks = logic(cmd, value, tasks)
}
