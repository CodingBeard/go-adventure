package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ConvertInput will convert a user's input into something the computer can read
type ConvertInput struct{}

//GetName will get the name
func (c *ConvertInput) GetName() string {
	return "Peach_Pie"
}

//Run will execute ConvertInput
func (c *ConvertInput) Run() error {
	// ask the user how many apples
	fmt.Print("How many times should I say 'Apple'?: ")
	reader := bufio.NewReader(os.Stdin)
	// read the user's input
	userInput, err := reader.ReadString('\n')
	// check if there was an error reading the input
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	// convert the user's input from a string to an int
	userInput = strings.TrimSuffix(userInput, "\r\n")

	input, err := strconv.ParseInt(userInput, 10, 64)

	// // do some error checking
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// loop and increment a counter until it reaches the number entered
	// by the user

	for c := int64(1); c <= input; c++ {

		// str1 := strconv.FormatInt(int64(input), 10)
		// str1 = "Apple"
		fmt.Println(fmt.Sprintf("%s %d", "Apple", c))
	}

	fmt.Println("Bitteschön")

	return nil

}
