package main

import (
	"container/ring"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/daodao97/egin/pkg/cache"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(time.Now().Unix())
	// api()
	// redis()
	// valid()
	// utils.ConsulKVTest()
	// utils.ConsulDeRegister()
	// channel()
	fmt.Println("over")
}

func channel() {
	heade := ring.New(5)
	for i := 0; i < 5; i++ {
		heade.Value = 0
		heade = heade.Next()
	}
	// heade.Value = 1
	// heade.Value = 2
	heade.Do(func(i interface{}) {
		fmt.Println(i)
	})
}

func redis() {
	cache.InitDb()

	redis := cache.Redis{Connection: "default"}

	val, err := redis.Get("key")

	fmt.Println(val, err)
}

func api() {
	for range time.Tick(time.Second * 2) {
		_, _ = exec.Command("bash", "-c", "ab -n 3000 -c 200 http://127.0.0.1:8080/user\\?check_in\\=2021-04-16\\&check_out\\=2021-04-20").Output()
		// for i := 0; i < 6; i++ {
		//	wg.Add(1)
		//	go func(index int) {
		//		res, err := lib.Get("http://127.0.0.1:8080/user", map[string]string{}, map[string]string{})
		//		if err != nil {
		//			log.Println("ERROR", err)
		//		} else {
		//			//log.Println("RESULT", res)
		//			fmt.Println(index, res.StatusCode == 200)
		//			defer res.Body.Close()
		//		}
		//		wg.Done()
		//	}(i)
		// }
		// wg.Wait()
	}
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
