package main // only one main per package

import "fmt" // for console

func main () {
	fmt.Print("Hello ")
	fmt.Println("World!")

	name := "Frico" // declare variable
	fmt.Println("Hello", name)
	fmt.Printf("Hello %s \n", name) // use %s for string

	fname := "Friko"
	lname := "Simon"
	fullName := fname + " " + lname // concatenate strings

	fmt.Println(fullName)

}