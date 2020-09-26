package task

import (
	"fmt"
	"os"
	"strconv"

	"github.com/divan/num2words"
	"github.com/rodaine/numwords"
)

//NumberConvert will convert numbers into words
type NumberConvert struct{}

//GetName gets the name
func (n *NumberConvert) GetName() string {
	return "Apple_Pie"
}

//GetDescr gets the description
func (n *NumberConvert) GetDescr() string {
	return "converts numbers into words"
}

//Run executes NumberConvert
func (n *NumberConvert) Run() error {
	wordToBeSaid, err := readTheInput("What should I say? :)")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	amountOfTimes, err := readTheInput(fmt.Sprintf("How many times should I say %s? :/", wordToBeSaid))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	amountOfTimes2, err := strconv.Atoi(amountOfTimes)
	//another error check
	if err != nil {
		amountOfTimes2, err = numwords.ParseInt(amountOfTimes)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
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
	return nil
}
