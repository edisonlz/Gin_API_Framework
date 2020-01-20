package main

import "fmt"

func main() {
    
    var slice1 []int = make([]int, 4)

    slice1[0] = 1
    slice1[1] = 2
    slice1[2] = 3
    slice1[3] = 4

    for index, value :=range slice1 {
        fmt.Printf("Slice at %d is: %d\n",index,value)
    }

}