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
type rectangle struct {
	length float64
	width  float64
}

//creating a func from a using a struct as an argument
func area(r rectangle) float64 {
	l := r.length
	w := r.width
	a := l * w
	return a
}
func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, "good morning")
}

func (s secretAgent) speak() {
	fmt.Println(s.firstName, s.lastName, "is hurt but alive")
}

// Method is differnt from a func, the syntax for mtd is
// func (r ReceiverType) funcName(parameters) (results)   example below

func main() {

	//calculating using the fnc and struct
	r1 := rectangle{20, 25}
	r2 := rectangle{30, 20}
	fmt.Println("the area of the 1st rectangle is: ", area(r1))
	fmt.Println("the area of the 2nd rectangle is: ", area(r2))

	if area(r1) > area(r2) {
		fmt.Println("The 1st rectangle has the largest area of ", area(r1))

	} else {
		fmt.Println("The 2nd rectangle has the largest area of ", area(r2))
	}
	// example of function
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
