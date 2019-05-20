package main

import (
	"fmt"
	"math"
	"errors"
)

// Use errors.New to construct a basic error message
func Sqrt(value float64) (float64, error){
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func main(){
	result, error := Sqrt(-1)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

	result1, error := Sqrt(9)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result1)
	}
}



