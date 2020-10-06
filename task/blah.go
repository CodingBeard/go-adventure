package task

import "fmt"

//TestTask is a task to be used for testing
type TestTask struct{}

//GetName will get the name
func (q *TestTask) GetName() string {
	return "Tomato_Pie"
}

//GetDescr will get the decription
func (q *TestTask) GetDescr() string {
	return "will be tested"
}

//Run will exectute TestTask
func (q *TestTask) Run() error {
	x := 2
	y := 3

	sum := x + y

	fmt.Println(sum)

	return nil
}
