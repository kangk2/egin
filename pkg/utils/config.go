package utils

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
    Address  string
    Mode     string
    Custom   interface{}
    Database Databases
    Redis    interface{}
    Logger   LoggerStruct
}

type Database struct {
Host    string
Port    int
User    string
Passwd  string
Options map[string]string
}

type Databases map[string]Database

type LoggerStruct struct {
    Type      string // stdout|file
    FileName  string
    Formatter string
    Level     int // 0 PanicLevel 5 InfoLevel 6 DebugLevel
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
    fmt.Println(Config)
    if err != nil {
        return
    }
}
