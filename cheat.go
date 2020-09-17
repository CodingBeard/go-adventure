package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func initializingAndUpdatingVariables() {
	// fmt.Println(myInitializedVariable) <- this will not work before we initialise the variable

	myInitializedVariable := "my initializing value"
	fmt.Println(myInitializedVariable) //this will print my initializing value

	myInitializedVariable = "my updated value"
	fmt.Println(myInitializedVariable) //this will print the updated value

	//myInitializedVariable := 0
	//reinitializing the variable will cause it to break
	//Changing the type of variable bzw. string to number will also break it

	//string := "myStringValue" <- using a reserved keyword as a varaible name will not work
	//reserved keywords will be a different color z.B. func, string, int, usw.
}
func slices() {
	//initializing the variable 'pizzaSlice' as an empty slice of integers(int)
	var pizzaSlice []int
	pizzaSlice = append(pizzaSlice, 10)
	fmt.Println(pizzaSlice) //the number '10' will will be printed in brackets
	pizzaSlice = append(pizzaSlice, 11)
	fmt.Println(pizzaSlice) // '10' and '11' will be together in the square brackets

	fmt.Println(pizzaSlice[1]) //just the number '11' will be printed

	pieSlice := []string{"abc", "def"}
	fmt.Println(pieSlice[0]) //will print 'abc'
	abc := pieSlice[0]
	// abc = "abc"
	// we want to convert 'b' to "b", from a character to a string
	fmt.Println(string(abc[1]))

	eggSlice := []string{"egg1", "egg2", "egg3"} //initializing 'eggSlice' as a string slice
	for position, egg := range eggSlice {        //ranging (iterating) over a list in the order we created it
		fmt.Println(position, egg) //printing the position of each egg in the slice
	}
}

func loops() {
	eggSlice := []string{"egg1", "egg2", "egg3"} //initializing 'eggSlice' as a string slice
	for position, egg := range eggSlice {        //ranging (iterating) over a list in the order we created it
		fmt.Println(position, egg) //printing the position of each egg in the slice
	}

	for position := range eggSlice {
		eggSlice[position] = "my " + eggSlice[position] //updating the value of eggs in the eggSlice
	}
	for position, egg := range eggSlice { //ranging (iterating) over a list in the order we created it
		fmt.Println(position, egg) //printing the position of each egg in the slice
	}

	for _, egg := range eggSlice { //using an underscore to ignore the first return value
		fmt.Println(egg)
		_ = egg //discards egg, just uses egg to make compiler happy
	}

	menge := "a, b, c"
	intTime := []int{} //variables are able to be used inside the scope below from outside, but variables within the scope will only
	//be read within the scope itself
	stringtime := strings.Split(menge, ",") //splits the list at a specific character to give a slice of strings
	for _, faden := range stringtime {
		fadenInt, err := strconv.Atoi(strings.TrimSpace(faden))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		intTime = append(intTime, fadenInt)
	}

	counter := 0
	for counter < 10 { //for every time the counter is lower than 10, run the code between the squiggly brackets
		counter++ // add one to counter
		fmt.Println(counter)
	}

	counter = 10
	for counter > 0 { //for every time the counter is larger than 0, run the code between the squiggly brackets
		counter-- // remove one from counter
		fmt.Println(counter)
	}

	for counter2 := 0; counter2 < 10; counter2++ { //same as before, just saving lines of code
		fmt.Println(counter2)
	}
}

func stringies() {
	// sillystring := []rune{'a', 'b', 'c'}
	sillystring := "abc"
	fmt.Println(string(sillystring[0])) //this will print the letter 'a' because it's in zeroth place
	fmt.Println(string(sillystring[2])) //this will print the letter 'c'
	for _, letter := range sillystring {
		fmt.Println(string(letter))
	}

	amountOfTimes2 := 1
	wordToBeSaid := "h\ni"
	i := 0 // 'i' is a standard counter variable
	for i < amountOfTimes2 {
		i++
		fmt.Println(fmt.Sprintf("%s number %d", wordToBeSaid, i)) //Sprintf allows variables to intermingle
		//%s, %d, usw are used as placeholders, followed by the proper type
	}
	multiLineString := `hi 
	moin
	moin` //back ticks allow for multi line strings without the use of "\n"

	fmt.Println(multiLineString)

	stringtime2 := "strings"
	intTime := "strings"
	if len(stringtime2) != len(intTime) { //len() measures the length of two strings
		fmt.Fprintln(os.Stderr, "stringtime2 and intTime not the same length") //if they are not the same length, print: (err)
		return                                                                 //if two strings are not the same length, it will throw an error and return
	}

	_ = stringtime2
	_ = intTime
}

func testFunction() { //just a test function. That's it.
	fmt.Println("Moin Moin.")
}

func returningFunction() string { //this function will return a sting every time it is called
	return "Hallöchen~"
}

func twoReturnFunction() (string, error) {
	return "Bitte wat?", errors.New("Plattdeutsch nicht erkannt") //returning two values bzw. string and error
}

func argueWithMe(arg1 int) (bool, error) { //bool = true or false

	if arg1 > 5 {
		return true, nil
	}
	return false, errors.New("The number you gave me was smaller than 5")
}

func addWithMe(integers ...int) (int, error) { //variadic argument can accept any number of integer argurments
	fmt.Println(integers)

	// todo: take integers and add them up, then return the result
	a := 0
	for _, i := range integers {
		a += i
	}
	return a, nil
}

func printSum(message string, integers ...int) {
	b := 1
	for _, i := range integers {
		b *= i
	}
	fmt.Println(message, b)
}

func functions() {
	testFunction() //'testFunction' does not accept any arg when calling it, also does not return anything
	//it will still execute that which is inside the function itself

	h := returningFunction()
	fmt.Println(h) //will print "Hallöchen~"

	p, err := twoReturnFunction()
	fmt.Println(p, err) //this will print the string and error from the twoReturnFunction

	a, err := argueWithMe(7)
	fmt.Println(a) //this would print the word 'true'; error will be 'nil'

	a, err = argueWithMe(2)
	fmt.Println(a) // this would print 'false'; error will be "The number you gave me...blahblahblah"

	addWithMe() //it printed nothing; it is okay to not give any arguments to a variadic

	addWithMe(2, 58, 7) //prints the three arguments

	c, err := addWithMe(2, 6, 10) //adds and prints the sum of the given integers
	fmt.Println(c)

	printSum("The sum is: ", 2, 10, 2) //this will print the sum of all of the integers prepended with the first string argument

}
