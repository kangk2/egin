package main

import (
    "fmt"
    "github.com/daodao97/egin/pkg/lib"
    "log"
    "sync"
)

var wg sync.WaitGroup

func main() {
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(index int) {
            res, err := lib.Get("http://127.0.0.1:8080/user", map[string]string{}, map[string]string{})
            if err != nil {
                log.Println("ERROR", err)
            } else {
                log.Println("RESULT", res)
                fmt.Println(index, res.StatusCode == 200)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
    fmt.Println("over")
}
