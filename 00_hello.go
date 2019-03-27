package main;

import ("fmt"
    "math"
    "math/rand")

    // comment
    /*
    comment
    */
func calculate() {
    fmt.Println("the square root of 6.25 is: ", math.Sqrt(6.25));
}

func main() {
    calculate();
    fmt.Println("number from 1 to 2", rand.Intn(2));
}
