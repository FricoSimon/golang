package main // only one main per package

import (
	"fmt" // for console
)

func multiply(x, y int) int {
	multiple := x * y
	return multiple
}

func names() (string, string) {
	return "Frico", "Simon"
}

// nested struct
type car struct {
	model string
	brand struct {
		name    string
		country string
	}
}
type rectangle struct {
	length, width int
}

// method without return value
func (r rectangle) area() {
	area := r.length * r.width
	fmt.Println("Area of rectangle:", area)
}

func main() {
	// call function
	fmt.Println(multiply(2, 3))

	name1, _ := names() // ignore second return value using dash
	fmt.Println(name1)

	// anonymous function
	func() { fmt.Println("I am an anonymous function") }() // call anonymous function

	// anonymous function with a name to a variable
	greeting := func() {
		fmt.Println("I am an anonymous function with a name / variable")
	}

	greeting()

	// parse float to int
	numberA := 10.5
	numberB := int(numberA)

	fmt.Println(numberB)

	// constant variable
	const pi = 3.14
	fmt.Printf("the value of pi is %f\n", pi)
	fmt.Printf("the type of pi is %T\n", pi)
	resultSprintf := fmt.Sprintf("the value of pi is %.2f and the type of pi is %T", pi, pi) // Sprintf returns a string
	fmt.Println(resultSprintf)

	// if statement with a short statement
	word := "Hello"
	//length := len(word)
	if length := len(word); length >= 5 { // scope of length is only in the if statement
		fmt.Println("The word has 5 characters")
	}

	// struct
	type people struct {
		name string
		age  int
	}

	friko := people{name: "Friko", age: 30}
	fmt.Println(friko)

	if friko.name == "Friko" {
		fmt.Println("correct name")
	} else {
		fmt.Println("wrong name")
	}

	//embedded struct
	type worker struct {
		people // take people struct
		job    string
	}

	engineer := worker{
		people: people{name: "Friko", age: 30},
		job:    "Engineer",
	}

	fmt.Println(engineer)
	fmt.Println(engineer.name) // access to embedded struct

	// nested struct call
	ferrari := car{model: "F40", brand: struct {
		name    string
		country string
	}{name: "Ferrari", country: "Italy"}}
	fmt.Println(ferrari)

	// method from struct call
	rect := rectangle{length: 10, width: 5}
	rect.area()
}
