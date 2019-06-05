/**
 * Structure
 */

package main

import(
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	bookId int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Programming"
	Book1.author = "Mike"
	Book1.subject = "GO Programming Tutorial"
	Book1.bookId = 100

	Book2.title = "PHP Programming"
	Book2.author = "Robert"
	Book2.subject = "PHP Programmig Tutorial"
	Book2.bookId = 200

	fmt.Printf(" Book 1 title is = %s\n", Book1.title)
	fmt.Printf(" Book 1 author is = %s\n", Book1.author)
	fmt.Printf(" Book 1 subject is = %s\n", Book1.subject)
	fmt.Printf(" Book 1 id is = %d\n", Book1.bookId)

	fmt.Printf(" Book 2 title is = %s\n", Book2.title)
	fmt.Printf(" Book 2 author is = %s\n", Book2.author)
	fmt.Printf(" Book 2 subject is = %s\n", Book2.subject)
	fmt.Printf(" Book 2 id is = %d\n", Book2.bookId)
}


/**
 * pointer
 */

// package main

// import(
// 	"fmt"
// )

// func main(){
// 	var a int = 20
// 	var p *int

// 	p = &a

// 	fmt.Printf(" Address of a variable = %x\n", &a)

// 	fmt.Printf(" Address stored in p variable = %x\n", p)

// 	fmt.Printf(" Value of *p variable = %d\n", *p)
// }



/**
 * Array and For Loop
 */

// package main

// import(
// 	"fmt"
// )

// func main(){
// 	var arr [10] int
// 	var i, j int

// 	for i = 0; i < 10; i++ {
// 		arr[i] = i + 100
// 	}

// 	for j = 0; j < 10; j++ {
// 		fmt.Printf(" Element[%d] = %d\n", j, arr[j])
// 	}
// }



/**
 * functins
 */

// package main

// import(
// 	"fmt"
// )

// func main() {
// 	var a int = 100
// 	var b int = 200
// 	var ret int

// 	ret = max(a, b)

// 	fmt.Printf("Max value is %d\n", ret)

// 	m, n := swap("Jim", "Umair")
// 	fmt.Println(m, n)
// }

// func max(num1, num2 int) int {
// 	var result int

// 	if (num1 > num2) {
// 		result = num1
// 	} else {
// 		result = num2
// 	}

// 	return result
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }


/**
 * Print Hello World
 */ 

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	fmt.Println("Hello WOrld!")
// }


/**
 * Error Handling
 */

// package main

// import (
// 	"fmt"
// 	"math"
// 	"errors"
// )

// // Use errors.New to construct a basic error message
// func Sqrt(value float64) (float64, error){
// 	if value < 0 {
// 		return 0, errors.New("Math: negative number passed to Sqrt")
// 	}
// 	return math.Sqrt(value), nil
// }

// func main(){
// 	result, error := Sqrt(-1)

// 	if error != nil {
// 		fmt.Println(error)
// 	} else {
// 		fmt.Println(result)
// 	}

// 	result1, error := Sqrt(9)

// 	if error != nil {
// 		fmt.Println(error)
// 	} else {
// 		fmt.Println(result1)
// 	}
// }


/**
 * Hello World
 */
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello Word")
// }



/**
 * Interface
 */
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Shape interface {
// 	area() float64
// }

// type Circle struct {
// 	x, y, radius float64
// }

// type Rectangle struct {
// 	width, height float64
// }

// func(circle Circle) area() float64 {
// 	return math.Pi * circle.radius * circle.radius
// } 

// func(rectangle Rectangle) area() float64 {
// 	return rectangle.width * rectangle.height
// }

// func getArea(shape Shape) float64 {
// 	return shape.area()
// }

// func main() {
// 	circle := Circle{x:0, y:0, radius:5}
// 	rectangle := Rectangle{width:5, height:10}

// 	fmt.Printf("Area of the circle %f\n", getArea(circle))
// 	fmt.Printf("Area of the rectangle %f\n", getArea(rectangle))
// }

