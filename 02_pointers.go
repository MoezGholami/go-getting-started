package main

import "fmt"

func swap(a, b *int) {
    var t int = *b
    *b = *a
    *a = t
}
func main() {
    x,y := 15,0
    a := &x
    *a = 19
    var b *int = &x
    fmt.Println(a)
    fmt.Println(*b)
    fmt.Println(*a)
    fmt.Println(x, y)
    swap(&x,&y)
    fmt.Println(x, y)
}
