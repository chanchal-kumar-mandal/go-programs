package main

import "fmt"

func main() {
    // Arithmetic Operators
    var m,n int = 100,25
    var a,b string = "John","Carry"
    var sum,difference,product,quotient,remainder int
    var sumstring string
    sum = m+n // +   sum   integers, floats, complex values, strings
    sumstring = a+" "+b
    difference = m-n // -   difference   integers, floats, complex values
    product = m*n // *   product   integers, floats, complex values
    quotient = m/n // /   quotient   integers, floats, complex values
    remainder = m%n //%   remainder  integers
    fmt.Println(sum)
    fmt.Println(sumstring)
    fmt.Println(difference)
    fmt.Println(product)
    fmt.Println(quotient)
    fmt.Println(remainder)

    // Bitwise Operators
    var x,y,z int
    x = 75
    y = 25
    z = x & y       // &    bitwise AND            integers
    fmt.Println(z)
    z = x | y       // |    bitwise OR             integers
    fmt.Println(z)
    z = x ^ y       // ^    bitwise XOR            integers
    fmt.Println(z)
    z = x &^ y      // &^   bit clear (AND NOT)    integers
    fmt.Println(z)

    // Comparison operators
    var num1 int = 50
    var num2 int = 60
    fmt.Println(num1==num2) // ==    equal
    fmt.Println(num1!=num2) // !=    not equal
    fmt.Println(num1<num2) <="" less="" fmt.println(num1<="num2)" or="" equal="" fmt.println(num1="">num2)  // >     greater
    fmt.Println(num1>=num2) // >=    greater or equal

    // Logical Operators
    var num1 int = 50
    var num2 int = 60
    
    if(num1!=num2 && num1<=num2){ // &&  Called Logical AND operator.
        fmt.Println("True")
    }
    
    if(num1!=num2 || num1<=num2){ // ||  Called Logical OR Operator
        fmt.Println("True")
    }
    
    if(!(num1==num2)){ // !  Called Logical NOT Operator. Use to reverses the logical state of its operand.
        fmt.Println("True")
    }

    // Assignment Operators
    var X int = 50
    var Y int = 60
    
    X+=Y    // += Add AND assignment operator
    fmt.Println(X)
    
    X=50
    Y=60
    X-=Y   // -= Subtract AND assignment operator
    fmt.Println(X)
    
    X=50
    Y=60
    X*=Y  // *= Multiply AND assignment operator
    fmt.Println(X)
    
    X=4
    Y=44
    X%=Y  // %= Modulus AND assignment operator
    fmt.Println(X)
    
    X=50
    Y=200
    X/=Y  // /= Divide AND assignment operator
    fmt.Println(X)
}