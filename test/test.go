package main

import (
    "fmt"
    "github.com/daodao97/egin/pkg/cache"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/go-playground/validator/v10"
    "log"
    "sync"
)

var wg sync.WaitGroup

func maina() {
    //api()
    //redis()
    valid()
    fmt.Println("over")
}

func redis() {
    cache.InitDb()

    redis := cache.Redis{Connection: "default"}

    val, err := redis.Get("key")

    fmt.Println(val, err)
}

func api() {
    for i := 0; i < 500; i++ {
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
}

func valid() {

    validate := validator.New()

    err := validate.Struct(1)
    processErr(err)

    err = validate.VarWithValue(1, 2, "eqfield")
    processErr(err)
}

func processErr(err error) {
    if err == nil {
        return
    }

    invalid, ok := err.(*validator.InvalidValidationError)
    if ok {
        fmt.Println("param error:", invalid)
        return
    }

    validationErrs := err.(validator.ValidationErrors)
    for _, validationErr := range validationErrs {
        fmt.Println(validationErr)
    }
}
