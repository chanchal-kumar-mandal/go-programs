package main

import "fmt"

func main() {
	
	var intSlice1 = make([]int,10)  // when length and capacity is same
	var intSlice2 = make([]int,10,20) // when length and capacity is different
	
	fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
	fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))
}