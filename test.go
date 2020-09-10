package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ask what should be said
	fmt.Println("What should I say? :) ")
	//read the user's input
	reader := bufio.NewReader(os.Stdin)
	wordToBeSaid, err := reader.ReadString('\n')
	//check for errors in the input
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	wordToBeSaid = strings.TrimSpace(wordToBeSaid)

	//ask how many times it should be said
	fmt.Println(fmt.Sprintf("How many times should I say %s? :/", wordToBeSaid))
	// fmt.Println("How many times should I say it? :/ ")
	//todo: update the second question with user input
	//todo: words into digits? z.B. "three" to "3"
	//read the user's input...again
	reader = bufio.NewReader(os.Stdin)
	amountOfTimes, err := reader.ReadString('\n')
	//error checking
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//convert the user input from string into an int
	amountOfTimes = strings.TrimSpace(amountOfTimes)
	amountOfTimes2, err := strconv.Atoi(amountOfTimes)
	//another error check
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//loop and increment a counter
	i := 0
	for i < amountOfTimes2 {
		i++
		fmt.Println(fmt.Sprintf("%s number %d", wordToBeSaid, i))
	}

	//use Sprintf with user's input from first question instead of literal word

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
