package main

import "fmt"

func main() {
    fmt.Println("hello world!")
    for i := 0; i < 3; i++ { // yes { is necessary even for a single statement
        fmt.Println(i)
    }
    for i := float64(0); i < 5; i+=2 { // var is not allowed either
        fmt.Println(i)
    }

    var outsideIndex int64 = 0
    for outsideIndex < 3 {
        fmt.Println(outsideIndex)
        outsideIndex++
    }

    for { // infinite loop
        outsideIndex += 10
        if outsideIndex > 10000 {
            break
        }
    }

     /*
    var i2 int32 = 1
    for {
        i2 += 1
        if i2 < 0 {
            break
        }
    }
    // */

    myArray := [...]string{"ali", "vali", "pirali"}
    for index, value := range myArray {
        fmt.Printf("%s at index %d\n", value, index)
    }
}
