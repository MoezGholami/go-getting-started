package main

import ("fmt")

func main() {
    fmt.Println("hello world")
    var grades map[string]float64 = make(map[string]float64)
    grades["gholam"] = 19
    grades["gholi"] = 20
    grades["ghamar"] = 18
    grades["moez"] = 8

    fmt.Println(grades["moez"])
    grades["moez"]+=1
    fmt.Println(grades["moez"])
    delete(grades,"moez")
    fmt.Println(grades["moez"])

    fmt.Println(grades)

    for key, value := range grades {
        fmt.Printf("the grade of %s is %f\n", key, value)
    }
}
