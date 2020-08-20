package model

import (
    "fmt"
    "github.com/daodao97/egin/pkg/db"
    "github.com/daodao97/egin/pkg/lib"
)

type UserEntity struct {
    Id       int  `json:"id"`
    Username string `json:"username"`
    RealName string `json:"realname"`
    Password string `json:"password"`
}

type UserM struct {
    db.BaseModel
    Entity UserEntity
}

func (m *UserM) Get(filter db.Filter, attr db.Attr) ([]UserEntity, error) {
    var result []UserEntity
    list, err := m.BaseModel.Get(filter, attr)
    if err != nil {
        return result, err
    }

    result = make([]UserEntity, len(list))

    for i, v := range list {
        tmp := make(map[string]interface{})
        for key, val := range v {
            tmp[key] = val
        }
        err := lib.UpdateStructByTagMap(&result[i], "json", tmp)
        fmt.Println(tmp, result[i])
        if err != nil {
            continue
        }
    }

    return result, nil
}

var UserModel UserM

func init() {
    entity := UserEntity{}
    UserModel = UserM{
        Entity: entity,
        BaseModel: db.BaseModel{
            Connection: "default",
            Table:      "user",
        },
    }
}
