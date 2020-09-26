package task

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//DoubleListCounter asks for lists and counts them
type DoubleListCounter struct{}

func (d *DoubleListCounter) countAsync(arg1 int, thing string, wg *sync.WaitGroup) {
	for i := 1; i <= arg1; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}

//GetName gets name
func (d *DoubleListCounter) GetName() string {
	return "Mango_Pie"
}

//GetDescr will get the description
func (d *DoubleListCounter) GetDescr() string {
	return "asks for lists and counts them"
}

//Run will execute the DoubleListCounter
func (d *DoubleListCounter) Run() error {
	// ask for a list of things to say
	response, err := readTheInput("Give me three words to say.")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// ask for a list of times to say each thing
	//todo: insert response into the question
	menge, err := readTheInput("How many times should I say each of those words?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// menge2, err := strconv.Atoi(menge)

	stringtime := strings.Split(menge, ",")
	//range over stringtime
	//todo: add to chat.go (inside outside explaination)
	intTime := []int{}
	for _, faden := range stringtime {
		fadenInt, err := strconv.Atoi(strings.TrimSpace(faden))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		intTime = append(intTime, fadenInt)
	}
	//convert the items in the string to an int
	//append to intTime slice with new int

	// use strings.Split() to split the lists on a specific character (E.G. split a, b, c on ',' to give you []string{"a", "b", "c"})
	stringtime2 := strings.Split(response, ",")
	//todo: add len() to cheat.go
	if len(stringtime2) != len(intTime) {
		fmt.Fprintln(os.Stderr, "stringtime2 and intTime not the same length")
		return errors.New("stringtime2 and intTime not the same length")
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(stringtime2))

	// range over each thing to say and spawn a go routine so they all say things in parallel threads
	//todo: add this to cheat.go (the use of places to combine two slices together)
	for places := range stringtime2 {

		// places = 0 then 1 then 2
		// intTime contains number 1 at place 0, number 2 at place 1, number 3 at place 2

		go d.countAsync(intTime[places], strings.TrimSpace(stringtime2[places]), wg)

	}

	wg.Wait()
	return nil
}
