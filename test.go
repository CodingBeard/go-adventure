package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/divan/num2words"
	"github.com/rodaine/numwords"
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
	// ask for a list of things to say
	response, err := readTheInput("Give me three words to say.")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// ask for a list of times to say each thing
	//todo: insert response into the question
	menge, err := readTheInput("How many times should I say each of those words?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
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
			return
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
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(stringtime2))

	// range over each thing to say and spawn a go routine so they all say things in parallel threads
	//todo: add this to cheat.go (the use of places to combine two slices together)
	for places := range stringtime2 {

		// places = 0 then 1 then 2
		// intTime contains number 1 at place 0, number 2 at place 1, number 3 at place 2

		go countAsync(intTime[places], strings.TrimSpace(stringtime2[places]), wg)

	}

	wg.Wait()

	fmt.Println(" :) ")
}

// wait for all the goroutines to finish

func mainWithPortals() {
	response, err := readTheInput("Do you think I can multitask?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if strings.TrimSpace(strings.ToLower(response)) == "yes" {
		fmt.Println("Well, thank you")
	} else {
		fmt.Println("Let me prove it then")
	}

	word1, err := readTheInput("What should I say first?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	amountOfWord1, err := readTheInput(fmt.Sprintf("And how many times should I say %s?", word1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	amountOfWordOne, err := strconv.Atoi(amountOfWord1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	word2, err := readTheInput("Great, now what should my second word be?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	amountOfWord2, err := readTheInput(fmt.Sprintf("And how many times should I say %s?", word2))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	amountOfWordTwo, err := strconv.Atoi(amountOfWord2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// let's pretend for the moment the go routines never run, what happens on each line here?
	wg := &sync.WaitGroup{}
	wg.Add(2) //the wg is waiting on two goroutines?

	go countAsync(amountOfWordOne, word1, wg)

	go func() {
		countSync(amountOfWordTwo, word2)
		wg.Done()
	}()

	wg.Wait() // waiting for the other to finish?

	fmt.Println("See?")

}

func appleQuestion() {
	wordToBeSaid, err := readTheInput("What should I say? :)")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	amountOfTimes, err := readTheInput(fmt.Sprintf("How many times should I say %s? :/", wordToBeSaid))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	amountOfTimes2, err := strconv.Atoi(amountOfTimes)
	//another error check
	if err != nil {
		amountOfTimes2, err = numwords.ParseInt(amountOfTimes)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	//loop and increment a counter
	//todo: type response into numbers as well
	i := 0
	for i < amountOfTimes2 {
		i++
		fmt.Println(
			fmt.Sprintf(
				"%s number %s",
				wordToBeSaid,
				num2words.Convert(i),
			),
		)
	}

	//use Sprintf with user's input from first question instead of literal word

}

func lookatme() {
	word := "thirty two"
	number := 0

	for i := 0; i < 1000; i++ {
		if num2words.Convert(i) == word {
			number = i
			fmt.Println(number)
			break
		}
	}

	fmt.Println(num2words.Convert(17))
	return
}

func appleit() {

	// ask the user how many apples
	fmt.Print("How many times should I say 'Apple'?: ")
	reader := bufio.NewReader(os.Stdin)
	// read the user's input
	userInput, err := reader.ReadString('\n')
	// check if there was an error reading the input
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// convert the user's input from a string to an int
	userInput = strings.TrimSuffix(userInput, "\r\n")

	input, err := strconv.ParseInt(userInput, 10, 64)

	// // do some error checking
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// loop and increment a counter until it reaches the number entered
	// by the user

	for c := int64(1); c <= input; c++ {

		// str1 := strconv.FormatInt(int64(input), 10)
		// str1 = "Apple"
		fmt.Println(fmt.Sprintf("%s %d", "Apple", c))
	}

	fmt.Println("Bitteschön")
}

func whatdoyouwantfromme() {
	fmt.Print("Is the world flat? [Yes/no]: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if strings.TrimSpace(strings.ToLower(userInput)) == "yes" {
		fmt.Println("Looooool")
	} else {
		fmt.Println("Richtiiiiig")
	}

	fmt.Println("You answered: ", userInput)
}

func oldTestingStuff() {
	x := 7
	y := 29
	sum := x + y

	_ = sum
	//fmt.Println(sum)

	// we are initialising a here by putting a : in front of the =
	a := []int{6, 8, 9, 10}
	// appendedToSlice = append(slice you want to append to, new element to go on the end of the slice)
	// we don't re-initialise it with another : or the compiler complains
	a = append(a, 23)

	counter := 0
	for counter < 10 {
		counter++
		// % E.G. 10 % 3 = 1
		a = append(a, 23+counter)
	}

	// use for loop to append numbers to a

	for key, value := range a {
		if value%2 == 0 {
			fmt.Println(key, value)
		} else {
			fmt.Println("Value is odd")
		}
	}

	// variables like staying the same type
	// a := 0 is the same as var a int, which is different to var a []
	b := 0

	for b < 10 {
		////fmt.Println(b)
		b++
	}

	fmt.Println("")
}

func otherMain() {
	var a []int
	a = append(a, 10)
	a = append(a, 11)
	a = append(a, 12)

	fmt.Println(a)

	// I typed "forr" and selected the range from the dropdown
	// The _ is an ignored variable, but it will be the key (position) in the slice, and the second variable before the := is the value

	{
		// squiggly brackets always come in pairs, if there's one on its own it will break things
	}
}
