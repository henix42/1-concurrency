package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Ёмкость 10, чтобы генератор успел заполнить send без блокировки на читателе.
	send := make(chan int, 10)
	receive := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			send <- rand.Intn(101)
		}
		close(send)
	}()

	go func() {
		for val := range send {
			receive <- val * val
		}
		close(receive)
	}()

	// Дожидаемся всех 10 чисел (канал закроется после квадратов) и выводим в консоль.
	for v := range receive {
		fmt.Println(v)
	}
	fmt.Println("Done")
}
