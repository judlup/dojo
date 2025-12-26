package main

import "testing"

func TestAddDeleteList(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Sample Task 1", Done: false},
		{Id: 2, Title: "Sample Task 2", Done: true},
	}
	tasks, err := ExecuteCommand("add", "New Task", tasks)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}

	tasks, err = ExecuteCommand("delete", "2", tasks)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks after deletion, got %d", len(tasks))
	}

	for _, task := range tasks {
		if task.Id == 2 {
			t.Errorf("Task with ID 2 should have been deleted")
		}
	}

	_, err = ExecuteCommand("delete", "5", tasks)
	if err == nil {
		t.Errorf("Expected error for non-existent task ID")
	}
	tasks, err = ExecuteCommand("list", "", tasks)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestEmptyTitle(t *testing.T) {
	tasks := []Task{}
	_, err := ExecuteCommand("add", "", tasks)
	if err == nil {
		t.Errorf("Expected error for empty task title")
	}
}

func TestDeleteInvalidID(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Sample Task 1", Done: false},
		{Id: 2, Title: "Sample Task 2", Done: true},
	}
	_, err := ExecuteCommand("delete", "abc", tasks)
	if err == nil {
		t.Errorf("Expected error for invalid task ID")
	}
}
