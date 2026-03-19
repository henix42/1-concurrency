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

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// func ping(url string, respCh chan <- int, errCh chan <- error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		errCh <- err
// 		return
// 	}
// 	respCh <- resp.StatusCode
// }

// func main() {
// 	path := flag.String("file", "url.txt", "path to url file")
// 	flag.Parse()
// 	file, err := os.ReadFile(*path)
// 	if err != nil {
// 		log.Fatalf("File reading error", err)
// 	}
// 	urlSlice := strings.Split(string(file), "\n")
// 	respCh:= make(chan int)
// 	errCh := make(chan error)

// 	for _, url := range urlSlice {
// 		go ping(url, respCh, errCh)
// 	}

// 	for range len(urlSlice) {
// 		select {
// 		case errRes := <- errCh:
// 			fmt.Println(errRes)
// 		case res := <- respCh:
// 			fmt.Println("Code #", res)

// 		}

// 	}

// }
