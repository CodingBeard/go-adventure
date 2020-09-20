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

func main() {
	app := task.Application{}
	app.AddTask(&task.DoubleListCounter{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.MultiTasking{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.NumberConvert{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.WordConvert{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.ConvertInput{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.RespondInput{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.Summing{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	app.AddTask(&task.FirstCode{})
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide a valid task name")
	}
	err := app.RunTask(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

}
