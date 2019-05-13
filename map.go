package main

import "fmt"

var employee = map[string]int{}
var employee1 = map[string]int{"Chanchal":1}
var employee2 = map[string]int{"Chanchal":1, "Kumar":2, "Mandal":3}

func main() {
	fmt.Println(employee)
	fmt.Println(employee1)
	fmt.Println(employee2)

	// using make() function for map
	var employee3 = make(map[string]int)
	employee3["Chanchal"] = 10
	employee3["Kumar"] = 20
	employee3["Mandal"] = 30
	fmt.Println(employee3)

	employee4 := make(map[string]int)
	employee4["Rohit"] = 100
	employee4["Hitman"] = 200
	employee4["Sharma"] = 300
	fmt.Println(employee4)

	// len() returns the number of element in a map
	fmt.Println(len(employee))
	fmt.Println(len(employee1))
	fmt.Println(len(employee4))

	// delete() function delets an element from given map associated with provided key
	delete(employee4, "Hitman")
	fmt.Println(employee4)
	fmt.Println(len(employee4))

	// Adding and Editting element in map
	var employee5 = map[string]int{"Virat":11, "Sachin":10}
	fmt.Println(employee5)

	employee5["Century"] = 165
	employee5["Rank"] = 1
	employee5["Virat"] = 18
	fmt.Println(employee5)

	// for range loop to iterate over a Map
	var employee6 = map[string]int{"Mark": 10, "Sandy": 20,
        "Rocky": 30, "Rajiv": 40, "Kate": 50}
    for key, element := range employee6 {
        fmt.Println("Key:", key, "=>", "Element:", element)
    }
}