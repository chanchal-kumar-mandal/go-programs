package main

import "fmt"

func main() {

	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
	
	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}
	
	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	for range "Hello" {
		fmt.Println("Hello")
	}

	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}