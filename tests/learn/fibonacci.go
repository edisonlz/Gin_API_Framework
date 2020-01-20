package main

import "fmt"

func main(){

    result := 0
    for i:=0;i<10;i++ {
        result = fibonacci(i)
        fmt.Printf("fibonacci(%d) is:%d\n",i,result)         
    }
}

func fibonacci(n int) (res int){

    if n<1{
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }

    return res
}


