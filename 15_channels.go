package main

import (
    "fmt"
    "time"
    "sync"
)

func asyncCalculation(c chan int, x, y int) {
    time.Sleep(time.Second)
    c <- x+y
}

func communicativeAdd(c chan int) {
    fmt.Println("inside communicativeAdd: waitin for values")
    x, y := <-c, <-c
    fmt.Println("inside communicativeAdd, got the values")
    c <- x+y
}

func synchronizedAddThroughChannel(c chan int, wg *sync.WaitGroup, x, y int) {
    defer wg.Done()
    time.Sleep(time.Second)
    c <- x+y
}

func main() {
    var addChannel chan int = make(chan int)
    communicativeChannel:= make(chan int)

    go asyncCalculation(addChannel, 1,2); go asyncCalculation(addChannel, 3,4) // these are parallel
    v1,v2 := <- addChannel, <- addChannel
    fmt.Println(v1,v2)

    for i := 0; i<2; i++ {
        go communicativeAdd(communicativeChannel)
        fmt.Println("yes, channel IO is blocking")
        time.Sleep(time.Second)
        communicativeChannel <- 2*i+1
        communicativeChannel <- 2*i+2
        fmt.Println("the result is: ", <- communicativeChannel)
    }

    N := 10
    bufferedChannel := make(chan int, N) /* buffer for N values
                         if it is anything less than N (say N-1 for example), we'll
                         have a deadlock as the last add will wait for the channel to have some space
                         (main thread should consume some value) and the main thread is waiting for the
                         last add to be done with wg
                         The question is how go monitors these dependencies on resources with different types
                         */
    var wg sync.WaitGroup
    for i := 0; i < N; i++ {
        wg.Add(1)
        go synchronizedAddThroughChannel(bufferedChannel, &wg, i, 0)
    }
    wg.Wait()
    close(bufferedChannel)
    fmt.Println("showing the serial results")
    for result := range bufferedChannel {
        fmt.Println(result)
    }
}
