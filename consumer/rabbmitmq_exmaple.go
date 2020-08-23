package main

import (
    "flag"
    "fmt"
    "log"

    "github.com/daodao97/egin/pkg/utils"
)

var mode = flag.String("mode", "", "运行模式")

func main() {
    flag.Parse()
    switch *mode {
    case "push":
        push()
    case "consume":
        consume()
    default:
        log.Fatal("place set mode")
    }
}

func push() {
    mq := utils.NewRabbitMQSimple("egin_amqp_test")
    err := mq.PublishSimple("this is a test message")
    if err == nil {
        fmt.Println("投递成功")
    } else {
        log.Fatalf("投递失败, 原因: %s", err)
    }
}

func consume() {
    mq := utils.NewRabbitMQSimple("egin_amqp_test")
    mq.ConsumeSimple(func(msg string) error {
        log.Printf("Received a message: %s", msg)
        return nil
    })
}
