package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
    send := make(chan int, 10)
    recieve := make(chan int, 10)
    var wg sync.WaitGroup

    // Горутина генерации чисел
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
    // Горутина возведения в квадрат
    go func() {
        defer wg.Done()
        for val := range send {
            recieve <- val * val
        }
        close(recieve)
    }()

    // main читает результаты
    for ch := range recieve {
        fmt.Println(ch)
    }

    wg.Wait()
    fmt.Println("Done")
}

