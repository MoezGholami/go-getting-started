package main

import (
    "time"
    "fmt"
)

func echo(s string) {
    for i := 0; i<3; i++ {
        fmt.Println(s)
        time.Sleep(time.Second)
    }
}

func main() {
    go echo("I am Gholam")
    go echo("I am Gholi")
    go echo("I am Ghamar")
    // if you comment the line below, the program will halt, that because main() ends
    // is there any joing command? Stay tuned with the rest of this code sample
    // Don't raise your expectations, it's going to be a terrible solution (not as neat as synchronized blocks in java)
    echo("I am Moez")
}
