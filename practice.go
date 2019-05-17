


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

