package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
    send := make(chan int)
    recieve := make(chan int)
    var wg sync.WaitGroup
    go func() {
        slice := make([]int, 10)
        for i := 0; i < 10; i++ {
            slice[i] = rand.Intn(101)    
        }
        for _, item := range slice {
            send <- item
        }
        close(send)
    }()
    wg.Add(1)
    go func() {
        defer wg.Done()
        for val := range send {
            recieve <- val * val
        }
        close(recieve)
    }()
    
    
    for ch := range recieve {
        fmt.Println(ch)
    }
    wg.Wait()
    
}

