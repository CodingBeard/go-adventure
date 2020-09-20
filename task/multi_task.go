package task

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//MultiTasking will allow several Tasks to run at the same time
type MultiTasking struct{}

func (m *MultiTasking) countSync(arg1 int, thing string) {
	for i := 1; i <= arg1; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func (m *MultiTasking) countAsync(arg1 int, thing string, wg *sync.WaitGroup) {
	for i := 1; i <= arg1; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}

//GetName will get the name
func (m *MultiTasking) GetName() string {
	return "Cherry_Pie"
}

//Run will execute MultiTasking
func (m *MultiTasking) Run() error {

	response, err := readTheInput("Do you think I can multitask?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	if strings.TrimSpace(strings.ToLower(response)) == "yes" {
		fmt.Println("Well, thank you")
	} else {
		fmt.Println("Let me prove it then")
	}

	word1, err := readTheInput("What should I say first?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	amountOfWord1, err := readTheInput(fmt.Sprintf("And how many times should I say %s?", word1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	amountOfWordOne, err := strconv.Atoi(amountOfWord1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	word2, err := readTheInput("Great, now what should my second word be?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	amountOfWord2, err := readTheInput(fmt.Sprintf("And how many times should I say %s?", word2))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	amountOfWordTwo, err := strconv.Atoi(amountOfWord2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// let's pretend for the moment the go routines never run, what happens on each line here?
	wg := &sync.WaitGroup{}
	wg.Add(2) //the wg is waiting on two goroutines?

	go m.countAsync(amountOfWordOne, word1, wg)

	go func() {
		m.countSync(amountOfWordTwo, word2)
		wg.Done()
	}()

	wg.Wait() // waiting for the other to finish?

	fmt.Println("See?")
	return nil
}
