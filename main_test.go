package main

import "testing"

func TestAddSum(t *testing.T) {
	if AddSum(3) != 5 {
		t.Error("Expected 3 + 2 to equal 5")
	}
}

func TableAddSum(t *testing.T) {
	var tests = []struct {
		input  int
		wanted int
	}{
		{3, 5},
		{-1, 1},
		{6, 8},
		{100, 102},
		{-5, -3},
	}

	for _, test := range tests {
		if output := AddSum(test.input); output != test.wanted {
			t.Error("Test failed:{}inputted, {} expected, recieved {}", test.input, test.wanted, output)
		}
	}
}
