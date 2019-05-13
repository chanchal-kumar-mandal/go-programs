package main

import (
	"fmt"
)

func main() {
	var s = "Japan"
	t := true
	if t {
		fmt.Println(s)
	}

	u := 100
	
	if u == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	x := 100
	
	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	if v := 100; v == 100 {
		fmt.Println("Germany")
	}
}