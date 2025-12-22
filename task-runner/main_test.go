package main

import "testing"

func TestAddDeleteList(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Sample Task 1", Done: false},
		{Id: 2, Title: "Sample Task 2", Done: true},
	}
	tasks, _ = logic("add", "New Task", tasks)
	
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}

	tasks, err := logic("delete", "2", tasks)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks after deletion, got %d", len(tasks))
	}
	_, err = logic("delete", "5", tasks)
	if err == nil {
		t.Errorf("Expected error for non-existent task ID")
	}
	tasks, err = logic("list", "", tasks)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestVoidTitle (t *testing.T) {
	tasks := []Task{}
	_, err := logic("add", "", tasks)
	if err == nil {
		t.Errorf("Expected error for empty task title")
	}
}

func TestDeleteInvalidID(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Sample Task 1", Done: false},
		{Id: 2, Title: "Sample Task 2", Done: true},
	}
	_, err := logic("delete", "abc", tasks)
	if err == nil {
		t.Errorf("Expected error for invalid task ID")
	}
}