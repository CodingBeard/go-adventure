package task

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//ErrUnknownTask is returned by RunTask when a task is unknown
var ErrUnknownTask error = errors.New("unknown task")

//A Task is the core unctionality of a task
type Task interface {
	GetName() string
	Run() error
}

//Application holds the tsks and allows it to execute
type Application struct {
	tasks []Task
}

//AddTask adds a task to the Application
func (a *Application) AddTask(task Task) {
	a.tasks = append(a.tasks, task)
}

//RunTask runs the task of a given name
func (a *Application) RunTask(name string) error {
	for _, task := range a.tasks {
		if task.GetName() == name {
			return task.Run()
		}
	}
	fmt.Println("Task canot be found.")
	return ErrUnknownTask
}

func readTheInput(question string) (string, error) {

	fmt.Println(question)

	reader := bufio.NewReader(os.Stdin)
	usersInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(usersInput), nil
}
