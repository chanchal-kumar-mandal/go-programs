package main

import (
    "fmt"
    "time"
    "sync"
)  

type testConcurrency struct {
    min int
    max int
    country string
}

func printCountry(test *testConcurrency, groupTest *sync.WaitGroup) {   
    for i :=test.max ; i>test.min; i-- { 
        time.Sleep(1*time.Millisecond)          
        fmt.Println(test.country)
    }

    fmt.Println()
    groupTest.Done()    
}

func  main() {

    // printCountry() // A normal function call that executes printCountry synchronously and waits for completing it
    // go printCountry() // A goroutine that executes printCountry asynchronously and doesn't wait for completing it

    groupTest := new(sync.WaitGroup)

    japan := new(testConcurrency)
    china := new(testConcurrency)
    india := new(testConcurrency)

    japan.country = "Japan"
    japan.min = 0
    japan.max = 5

    china.country = "China"
    china.min = 0
    china.max = 5

    india.country = "India"
    india.min = 0
    india.max = 5

    go printCountry(japan, groupTest)
    go printCountry(china, groupTest)
    go printCountry(india, groupTest)

    groupTest.Add(3) // The Add method is used to add a counter to the WaitGroup
    groupTest.Wait() // Wait blocks the execution of the program until the WaitGroup counter is zero
}