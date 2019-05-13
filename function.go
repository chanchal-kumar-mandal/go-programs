package main

import "fmt"

func main() {
	fmt.Println(add(20, 30))
}

// Return types can have names
func add(x int, y int) (total int) {
	total = x + y
	return total
}