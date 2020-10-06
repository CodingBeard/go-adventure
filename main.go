package main

import (
	"bufio"
	"fmt"
	"go-adventure/task"
	"os"
	"strings"
	"sync"
	"time"
)

func readTheInput(question string) (string, error) {

	fmt.Println(question)

	reader := bufio.NewReader(os.Stdin)
	usersInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(usersInput), nil
}

func countSync(arg1 int, thing string) {
	for i := 1; i <= arg1; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countAsync(arg1 int, thing string, wg *sync.WaitGroup) {
	for i := 1; i <= arg1; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}

func addSum(x int) (result int) {
	result = x + 2
	return (result)
}

func main() {

	result := addSum(3)
	fmt.Println(result)

	app := task.Application{}
	app.AddTask(&task.DoubleListCounter{})

	app.AddTask(&task.MultiTasking{})

	app.AddTask(&task.NumberConvert{})

	app.AddTask(&task.WordConvert{})

	app.AddTask(&task.ConvertInput{})

	app.AddTask(&task.RespondInput{})

	app.AddTask(&task.Summing{})

	app.AddTask(&task.FirstCode{})

	app.AddTask((&task.ReadWrite{}))

	taskName := ""
	if len(os.Args) > 1 {
		taskName = os.Args[1]
	}
	err := app.RunTask(taskName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	return

}
