package task

import "fmt"

//FirstCode appends numbers to a slice
type FirstCode struct{}

//GetName will get the name
func (f *FirstCode) GetName() string {
	return "Strawberry_Pie"
}

//GetDescr gets the description
func (f *FirstCode) GetDescr() string {
	return "appends numbers to a slice"
}

//Run will execute FirstCode
func (f *FirstCode) Run() error {
	var a []int
	a = append(a, 10)
	a = append(a, 11)
	a = append(a, 12)

	fmt.Println(a)

	return nil
}
