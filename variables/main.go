package main

import "fmt"

var x = 2

//or var x int (used to declare an int without assignng. its value is 0 when not assigned)
// x := 5 is a shorter way but this only works inside a func body

func main() {
	x := 5
	fmt.Println(x)
	fmt.Printf("%T \n", x) // %T for type of variable
	// here the output is 5
}
