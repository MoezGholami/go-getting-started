package main

import (
    "time"
    "fmt"
    "sync"
)


func cleanUp() {
    if r := recover(); r != nil {
        fmt.Println("inside recovery", r)
    }
}

func raiseNilError() {
    var x *int = nil
    *x++
}

func echo(s string, wg *sync.WaitGroup) {
    defer wg.Done()
    defer cleanUp()
    defer fmt.Println("done with echo again, LIFO order")
    defer fmt.Println("done with echo")
    if s == "I am Gholi" {
        panic("no Gholi allowed")
    } else if s == "I am Gholam" {
        raiseNilError()
    }
    for i := 0; i<3; i++ {
        fmt.Println(s)
        time.Sleep(time.Second)
    }
}

func main() {
    var wg sync.WaitGroup
    var messages = [...]string {"I am Gholam" , "I am Gholi" , "I am Ghamar"}
    for _, m := range messages {
        wg.Add(1)
        go echo(m, &wg)
    }
    wg.Wait()
}
