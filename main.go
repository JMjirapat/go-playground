package main

import (
	"fmt"

	"github.com/JMjirapat/go-playground/function"
	referencetype "github.com/JMjirapat/go-playground/reference-type"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

// In go, scope is defined by the package
// If a variable, function, or type starts with a capital letter, it is exported
// If a variable, function, or type starts with a lowercase letter, it is private
// In go, different file but same package can access private variables, functions, and types

func main() {
	// variable declaration
	// const <variable_name> <data_type> = <value>
	const pi float64 = 3.14159265359
	// In go, we can declare variables in 3 ways. The variables are initialized to default values
	// 1) var <variable_name> <data_type>
	var x int

	// 2) var <variable_name> <data_type>? = <value>
	var y = 10
	var y2 int = 10

	// 3) <variable_name> := <value>
	// This is a shorthand for declaring and initializing a variable
	// The type of the variable is inferred from the value
	// This can only be used inside a function
	z := 10

	fmt.Println(x, y, y2, z, pi)

	// multiple variable declaration
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// variable type conversion
	var i int = 10
	var j float64 = float64(i)
	fmt.Println(j)

	fmt.Println(function.Add(1, 2))
	fmt.Println(function.Add3(1, 2, 3))
	fmt.Println(function.AddArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(function.AddVariadic(1, 2, 3, 4, 5))

	// defer
	// defer is used to delay the execution of a function until the surrounding function returns

	defer fmt.Println("World")
	fmt.Println("Hello")

	// panic and recover
	// panic is used to terminate the program abruptly
	// recover is used to regain control of a panicking goroutine
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	referencetype.MakeForReference()

	panic("Panic")

}
