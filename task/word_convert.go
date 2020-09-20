package task

import (
	"fmt"

	"github.com/divan/num2words"
)

//WordConvert will convert written numbers into ints
type WordConvert struct{}

//GetName will get the name
func (w *WordConvert) GetName() string {
	return "Banana_Pie"
}

//Run will execute WordConvert
func (w *WordConvert) Run() error {
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
	return nil
}
