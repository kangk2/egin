package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/daodao97/egin/pkg/lib"
)

type ConfigStruct struct {
	Name     string
	Address  string
	Mode     string
	Custom   interface{}
	Database Databases
	Redis    map[string]Redis
	Logger   LoggerStruct
	Lan      string
	Auth     struct {
		Cors struct {
			Enable           bool
			AllowOrigins     []string // 允许源列表
			AllowMethods     []string // 允许的方法列表
			AllowHeaders     []string // 允许的头部信息
			AllowCredentials bool     // 允许暴露请求的响应
		}
		IpAuth struct {
			Enable        bool
			AllowedIpList []string
		}
	}
	Jwt struct {
		Secret      string
		TokenExpire int64
	}
	RabbitMQ map[string]Rabbitmq
	Consul   string
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

type Redis struct {
	Host     string
	Port     int
	DB       int
	Password string
}

type Rabbitmq struct {
	Host   string
	Port   int
	User   string
	Passwd string
	Vhost  string
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

	kv, err := ConsulKV(Config.Consul)
	if err != nil {
		return
	}
	// 远程配置只能回覆盖式的, 不支持删除某个配置
	remoteConfKey := fmt.Sprintf("%s/%s", Config.Name, Config.Mode)
	kp, _, err := kv.Get(remoteConfKey, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(kp.Value, &Config)
	if err != nil {
		return
	}
	go func() {
		for range time.Tick(time.Second * 2) {
			kp, _, err := kv.Get(remoteConfKey, nil)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = json.Unmarshal(kp.Value, &Config)
			fmt.Println(Config)
		}
	}()
}
