package main

import (
	"fmt"
)

var m = "Mexico"

func main() {
	var i int = 10
	var s string = "Japan"
	fmt.Println(i)
	fmt.Println(s)

	// Assignment after initialization
	var intVar int
	var strVar string
	
	intVar = 10
	strVar = "Australia"
	
	fmt.Println(intVar)
	fmt.Println(strVar)

	// Short Variable Declaration
	s := "Japan"
	fmt.Println(s)

	// Scope of Variables Defined by Brace Brackets
	fmt.Println(s)
	x := true
	
	if x {
		y := 1
		if x != false {
			fmt.Println(m)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}