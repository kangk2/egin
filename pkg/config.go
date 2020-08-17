package pkg

import (
    "encoding/json"
    "fmt"
    "github.com/joho/godotenv"
    "io/ioutil"
    "os"
    "regexp"
    "strings"
    "github.com/daodao97/egin/pkg/lib"
)

type ConfigStruct struct {
    Address string
    Mode    string
}

var Config ConfigStruct

func init() {
    if err := godotenv.Load(".env"); err != nil {
        fmt.Printf("not found .env %s", err)
    }

    data, err := ioutil.ReadFile("./conf/app.json")

    str := string(data)

    re, _ := regexp.Compile("<.*>")

    all := re.FindAllString(str, -1)

    for i := range all {
        s := all[i]
        factory := lib.String{Str: s}
        r := os.Getenv(factory.TrimLeft("<").TrimRight(">").Done())
        str = strings.Replace(str, s, r, -1)
    }

    if err != nil {
        fmt.Printf("load config fail %s", err)
    }

    err = json.Unmarshal([]byte(str), &Config)
    if err != nil {
        return
    }
}