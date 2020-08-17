package pkg

import (
    "encoding/json"
    "fmt"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/joho/godotenv"
    "io/ioutil"
    "os"
    "regexp"
    "strings"
)

type ConfigStruct struct {
    Address string
    Mode    string
}

var Config ConfigStruct

func init() {
    if err := godotenv.Load(".env"); err != nil {
        fmt.Printf("not found .env\n%s\n", err)
    }

    data, err := ioutil.ReadFile("./config/app.json")

    if err != nil {
        fmt.Printf("load config/app.json fail\n%s\n", err)
        os.Exit(2)
    }

    str := string(data)

    re, _ := regexp.Compile("<.*>")

    all := re.FindAllString(str, -1)

    for i := range all {
        s := all[i]
        factory := lib.String{Str: s}
        r := os.Getenv(factory.TrimLeft("<").TrimRight(">").Done())
        str = strings.Replace(str, s, r, -1)
    }

    err = json.Unmarshal([]byte(str), &Config)
    if err != nil {
        return
    }
}
