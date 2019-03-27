package main

import (
    "fmt"
)

type functionInputError struct {
    input int
    functionName string
}

func (fe *functionInputError) Error() string {
    return fmt.Sprintf("the arguement %d is invalid for function %s", fe.input, fe.functionName)
}

func fib(n int) (int, error) {
    if n > 10000 {
        panic("too big to calculate fib for n > 10000")
    } else if n < 0 {
        return -1, &functionInputError{n, "fib"}
    } else if n < 2 {
        return n, nil
    } else {
        fn, fn1, fn2 := 1, 1, 0
        for i := 1; i < n; i++ {
            fn = fn1 + fn2
            fn2 = fn1
            fn1 = fn
        }
        return fn, nil
    }
}

func main() {
    for i := -2; i < 10; i++ {
        f, e := fib(i)
        fmt.Println("fib of", i, "is", f, "with error", e)
    }
}
