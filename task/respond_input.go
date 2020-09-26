package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//RespondInput will respond according to a user's input
type RespondInput struct{}

//GetName will get the name
func (r *RespondInput) GetName() string {
	return "Blueberry_Pie"
}

//GetDescr will get the description
func (r *RespondInput) GetDescr() string {
	return "responds according to a user's input"
}

//Run will execute RespondInput
func (r *RespondInput) Run() error {
	fmt.Print("Is the world flat? [Yes/no]: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	if strings.TrimSpace(strings.ToLower(userInput)) == "yes" {
		fmt.Println("Looooool")
	} else {
		fmt.Println("Richtiiiiig")
	}

	fmt.Println("You answered: ", userInput)

	return nil
}
