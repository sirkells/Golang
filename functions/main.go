package main

import "fmt"

//Functions
// format of declaring:   func(receiver) identifier(parameters)(returns){<code>}

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	licensetokill bool
}

// example of function

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, "good morning")
}

func (s secretAgent) speak() {
	fmt.Println(s.firstName, s.lastName, "is hurt but alive")
}
func main() {

	p1 := person{
		"Kelechi",
		"Igbokwe",
		31,
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
			34,
		},
		true,
	}
	sa1.speak()
	p1.speak()

}
