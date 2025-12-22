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

func logic(cmd string, value string, tasks []Task) ([]Task, error) {
	id := len(tasks) + 1
	switch cmd {
	case "add":
		task := Task{
			Id:    id,
			Title: value,
			Done:  false,
		}
		if value == "" {
			return nil, fmt.Errorf("task title cannot be empty")
		}
		tasks = append(tasks, task)
		return tasks, nil		
	case "list":
		// log("Todo List Tasks")
		// for _, task := range tasks {
		// 	log(fmt.Sprintf("Task ID: %d, Title: %s, Done: %v", task.Id, task.Title, task.Done))
		// }
		return tasks, nil
	case "delete":
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
	default:
		return tasks, fmt.Errorf("unknown command: %s", cmd)
	}
	return tasks, nil
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
	tasks, err := logic(cmd, value, tasks)
	if err != nil {
		log(err.Error())
	} else {
		log("Operation completed successfully")
		fmt.Printf("%+v\n", tasks)
	}
}
