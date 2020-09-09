package main

import (
	"fmt"
	"strconv"
)

func whatif() {
	a := "Apple"
	_, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Error converting str to int:")
	}
}
