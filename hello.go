package main // only one main per package

import "fmt" // for console

func main () {
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

}