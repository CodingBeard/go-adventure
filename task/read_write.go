package task

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//ReadWrite will read and write to a file
type ReadWrite struct{}

//GetName will get the name
func (r *ReadWrite) GetName() string {
	return "Orange_Pie"
}

//GetDescr will get the description
func (r *ReadWrite) GetDescr() string {
	return "reads and writes to a file"
}

//AppendFile will append the user's input to the given file
func AppendFile(s string) {
	file, err := os.OpenFile("file_read_write.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	words, err := file.WriteString(s)
	if err != nil {
		log.Fatalf("failed to write to file: %s", err)
	}
	_ = words
	fmt.Println("Saved!")
}

//Run will execute ReadWrite
func (r *ReadWrite) Run() error {
	data, err := ioutil.ReadFile("file_read_write.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return err
	}
	fmt.Println("Here it is:", string(data))

	response, err := readTheInput("Shall we add more text to the file?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	if strings.TrimSpace(strings.ToLower(response)) == "yes" {
		addedText, err := readTheInput("What would we like to add to the file?")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		AppendFile(addedText)

	} else {
		fmt.Println("Okay")
	}

	return err
}
