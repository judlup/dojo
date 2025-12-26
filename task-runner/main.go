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

type Command func(tasks []Task, value string) ([]Task, error)

func AddTask(tasks []Task, title string) ([]Task, error) {
	id := len(tasks) + 1
	if title == "" {
		return tasks, fmt.Errorf("task title cannot be empty")
	}
	task := Task{
		Id:    id,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	return tasks, nil
}

func ListTasks(tasks []Task, value string) ([]Task, error) {
	return tasks, nil
}

func DeleteTask(tasks []Task, value string) ([]Task, error) {
	taskId, err := strconv.Atoi(value)
	found := false
	if err != nil {
		return tasks, fmt.Errorf("invalid task ID")
	}
	for i, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return tasks, fmt.Errorf("task ID %d not found", taskId)
	}
	return tasks, nil
}

var commands = map[string]Command{
	"add":    AddTask,
	"list":   ListTasks,
	"delete": DeleteTask,
}

func ExecuteCommand(cmd string, value string, tasks []Task) ([]Task, error) {
	if commandFunc, exists := commands[cmd]; exists {
		return commandFunc(tasks, value)
	}
	return tasks, fmt.Errorf("unknown command: %s", cmd)
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
	// tasks, err := logic(cmd, value, tasks)
	tasks, err := ExecuteCommand(cmd, value, tasks)
	if err != nil {
		log(err.Error())
	} else {
		log("Operation completed successfully")
		fmt.Printf("%+v\n", tasks)
	}
}
