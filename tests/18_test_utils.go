package tests

func Add(x, y int) int {
    return x+y
}

func Divide(x, y int) int {
    if y == 0 {
        panic("custom divide error")
    } else {
        return x/y
    }
}
