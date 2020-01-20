package main

//go-in-practice.pdf
//https://books.studygolang.com/gobyexample/
//https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md
//go 中文
//https://go.fdos.me/07.2.html


import ("fmt"
        "time")


func sum(a []int, c  chan int) {
    total := 0

    for _,v :=range a {
        total += v
    }

    c <- total
}




func sum_main(){

    a := []int{1,2,3,4,5,6,7,8}
    c := make(chan int)

    go sum(a[:len(a)/2],c)
    go sum(a[len(a)/2:],c)

    x, y := <-c ,<-c

    fmt.Println("sum=",x+y , "half=",x,y)
}

func unblock_main(){
    c := make(chan int, 1)
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)
}


/* fab */
func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x + y
    }
    close(c)
}


func fab_main(){

    c :=make(chan int, 10)

    go fibonacci(cap(c), c)

    for i := range c {
        fmt.Println(i)
    }
}

/* select route */

//Select
//我们上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢
//，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。

func select_fibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
        case c <- x:
            x, y = y, x + y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}


func select_fab_main() {

    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    select_fibonacci(c, quit)
}

/* Timeout */

func timeout_main() {

    c := make(chan int)
    o := make(chan bool)

    go func() {
        for {
            select {
                case v := <- c:
                    println(v)
                case <- time.After(5 * time.Second):
                    println("timeout")
                    o <- true
                    break
            }
        }
    }()

    <- o
}

func main() {

    sum_main()
    //unblock_main()
    //fab_main()
    //select_fab_main()
    //timeout_main()
    
}