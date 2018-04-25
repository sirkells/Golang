//"A method is a function with an implicit first argument, called a receiver."
//Syntax of method.
//func (r ReceiverType) funcName(parameters) (results)
//Let's change our example using method instead.

package main

import (
    "fmt"
    "math"
)

type circle struct {
    radius float64
}

type rectangle struct {
    width, height float64
}
type human struct {
    name  string
    age   int
    phone string
}

type student struct {
    human  // anonymous field
    school string
}

type employee struct {
    human
    company string
}

type robot struct {
	name  string
    age   int
	job string
	
}

type security struct {
	robot
	company string
}

// define a method in Human
func (h human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
// this mtd is now a field of human
// mtd for robot
func (r robot) RobHi() {
    fmt.Printf("Hi, I am %s and i am a robot in charge of %s for %d years now\n", r.name, r.job, r.age)
}
//method
func (c circle) Area() float64 {
    return c.radius * c.radius * math.Pi
}

// method
func (r rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
//Inheritance of method
// If an anonymous field has methods, then the struct that contains the field will have all the methods from it as well.
// here the struct employee and student inherited the func SayHi from fthe mtd in human
	sam := employee{human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := student{human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	john := security{robot{"Bingo", 4, "securing the building"}, "Kongtron Tech"}

    sam.SayHi()
	mark.SayHi()
	// here the struct security  inherited the func RobHi from fthe mtd in robot
	john.RobHi()
	
	
    c1 := circle{10}
    c2 := circle{25}
    r1 := rectangle{9, 4}
    r2 := rectangle{12, 2}

    fmt.Println("Area of c1 is: ", c1.Area())
    fmt.Println("Area of c2 is: ", c2.Area())
    fmt.Println("Area of r1 is: ", r1.Area())
    fmt.Println("Area of r2 is: ", r2.Area())
}

//In the example above, the Area() methods belong to both Rectangle and Circle respectively, so the receivers are Rectangle and Circle.there are diff customisable types
//Examples of customized types:
type age int
type money float32
type months map[string]int

//m := months {
    //"January":31,
    //"February":28,
    //...
    //"December":31,
//}