package model

import (
    "github.com/daodao97/egin/pkg/db"
)

type UserEntity struct {
    Id       int64  `json:"id"`
    Username string `json:"name"`
    Realname string `json:"realname"`
}

var UserModel db.BaseModel

func init() {
    entity := UserEntity{}
    UserModel = db.BaseModel{
        Connection: "default",
        Table:      "user",
        Entity:     entity,
    }
}
