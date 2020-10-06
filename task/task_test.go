package task

import (
	"go-adventure/task"
	"testing"
)

//Testss will test the task
type Testss struct {
	x   int
	y   int
	sum int
}

var testResults = []Testss{
	{2, 3, 5},
}

func TestTestss(t *testing.T) {
	for _, test := range testResults {
		taskToTest := &task.TestTask{}
		result := taskToTest(test.x, test.y)
		if result != test.sum {
			t.Fatal("Expected Result: {}", test.sum)
		}
	}
	return
}
