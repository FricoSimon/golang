package main // only one main per package

import "fmt" // for console

// function signature
func sum(x, y int) int {
		return x + y
	} 

func multiply (x, y int) int {
		multiple := x * y
		return multiple
	}

func main () {
	// call function
	fmt.Println(sum(2, 3))

	// call function
	fmt.Println(multiply(2, 3))

	// print to console
	fmt.Print("Hello ")
	fmt.Println("World!")

	// string interpolation
	name := "Frico" // declare variable
	fmt.Println("Hello", name)
	fmt.Printf("Hello %s \n", name) // use %s for string

	// concatenate strings
	fname := "Friko"
	lname := "Simon"
	fullName := fname + " " + lname 

	fmt.Println(fullName)

	// sum of two numbers
	a := 10
	b := 2
	fmt.Println(a + b) 

	// anonymous function
	func() {
		fmt.Println("I am an anonymous function")
	}()
	
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
	} else {
		fmt.Println("The word does not have 5 characters")
	} 
	
}