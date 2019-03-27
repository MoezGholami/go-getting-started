package main;

import ("fmt"
    )

func add(x float64, y float64) float64 {
    return x+y
}

func add2(x, y float64) float64 {
    return x+y
}

func add3(x, y float32) float32 {
    return x+y
}

func echo(x, y int) (int, int) {
    return x,y
}

const pi float64 = 3.141592654

func multipleLength(a,b string) (int, int) {
    return len(a), len(b)
}

func main() {
    fmt.Println(echo(1,2))
    fmt.Println("sum of 1 and 2.3 is: ", add(1,2.3));
    var a float64 = 5.6
    var b float64 = 9.5

    fmt.Println("sum of", a, "and", b, "is :", add2(a,b))

    var c, d float64 = -17.3, 19
    fmt.Println("sum of", c, "and", d, "is :", add2(c,d))

    e,f := 1.0,2
    fmt.Println("sum of", e, "and", f, "is :", add3(float32(e),float32(f)))

    fmt.Println(pi);

    fmt.Println(multipleLength("Ali", "Moezgholami"))
}
