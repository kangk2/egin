package utils

import (
    "encoding/json"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/joho/godotenv"
    "io/ioutil"
    "log"
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
    Lan      string
    Auth struct{
        IpAuth struct{
            Enable bool
            AllowedIpList []string
        }
    }
}

type Database struct {
    Host     string
    Port     int
    User     string
    Passwd   string
    Database string
    Driver   string
    Options  map[string]string
    Pool     struct {
        MaxOpenConns int
        MaxIdleConns int
    }
}

type Databases map[string]Database

type LoggerStruct struct {
    Type      string // stdout|file
    FileName  string
    Formatter string
    Level     int // 0 PanicLevel 5 InfoLevel 6 DebugLevel
}

var Config ConfigStruct

var defaultConfig = "{\"address\":\"127.0.0.1:8080\",\"mode\":\"debug\",\"custom\":[1,2,3],\"logger\":{\"type\":\"stdout\",\"fileName\":\"tmp/app.log\",\"level\":5},\"database\":{\"default\":{\"host\":\"localhost\",\"port\":3306,\"user\":\"root\",\"passwd\":\"root\",\"database\":\"hyperf_admin\",\"pool\":{\"maxOpenConns\":10,\"maxIdleConns\":5},\"options\":{}}}}"

func init() {
    if err := godotenv.Load(".env"); err != nil {
        log.Printf("load .env fail: %s", err)
    }

    // TODO 支持命令行参数
    data, err := ioutil.ReadFile("app.json")

    if err != nil {
        log.Printf("load app.json fail: %s, will use default config", err)
    }

    str := string(data)

    if str == "" {
        str = defaultConfig
    }

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
