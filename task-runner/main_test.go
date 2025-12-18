package main

import "testing"

func TestLogin(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Sample Task 1", Done: false},
		{Id: 2, Title: "Sample Task 2", Done: true},
	}
	tasks = logic("add", "New Task", tasks)
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
		tasks = logic("delete", "1", tasks)
		if len(tasks) != 2 {
			t.Errorf("Expected 2 tasks after deletion, got %d", len(tasks))
			tasks = logic("list", "", tasks)
			if len(tasks) != 2 {
				t.Errorf("Expected 2 tasks after listing, got %d", len(tasks))
			}
		}
	}

}
