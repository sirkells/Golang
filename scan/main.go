package main

import "fmt"

var x int



func main() {
	fmt.Println("please enter a number: ")
	fmt.Scan(&x)
	fmt.Println(x)
	fmt.Printf("%T \n", x) // %T for type of variable
	// here the output is 5
}