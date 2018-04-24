package main

import "fmt"

// there are three data structure types namesly Slice(list), Map(Dictionary) and Struct(classes)

// this is a struct
type person struct {
	firstName string
	lastName  string
	age       int
}

type rectangle struct {
	length float64
	width float64

}

func area(r rectangle) float64 {
	l := r.length
	w := r.width
	a := l * w
	return a

}


// defining an array
var arr [5] int // arr is name of array, [] is the length of the array. type is the type 
//defining a slice: slice is used instead of an array when you dont know how long the arry would be
var slice [] int 
// or slice := [] int{1,2,3,4,5}

func main() {
	r1 := rectangle{20, 25}
	r2 := rectangle{30, 20}
	fmt.Println("the area of the 1st rectangle is: ", area(r1))
	fmt.Println("the area of the 2nd rectangle is: ", area(r2))
	

	if area(r1) > area(r2) {
		fmt.Println("The 1st rectangle has the largest area of ", area(r1))
		
	} else {
		fmt.Println("The 2nd rectangle has the largest area of ", area(r2))
	}
	
	//SLICE
	x := []int{1, 2, 3, 4, 5} // this is called slice. slice is a list. this is how you create a slice of numbers
	fmt.Println(x)
	// MAP
	m := map[string]int{
		"John": 20,
		"Mike": 34,
	}
	fmt.Println(m)

	p := person{
		"Kelechi",
		"Igbokwe",
		31,
	}
	fmt.Println(p)

}
