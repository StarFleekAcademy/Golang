// https://gobyexample.com/values
// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
package main

import "fmt"

func main() {
	// Strings, which can be added together with +.
	fmt.Println("Go" + "lang")
	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	// Booleans, with boolean operators as you’d expect.
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
