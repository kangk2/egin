package model

import (
    "github.com/daodao97/egin/pkg/db"
    "github.com/daodao97/egin/pkg/lib"
)

type UserEntity struct {
    Id       int    `json:"id"`
    Username string `json:"username"`
    RealName string `json:"realname"`
    Password string `json:"password"`
}

type UserModel struct {
    db.BaseModel
}

func (m *UserModel) Get(filter db.Filter, attr db.Attr) ([]UserEntity, error) {
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
        if err != nil {
            continue
        }
    }

    return result, nil
}

var User UserModel

func init() {
    User = UserModel{
        BaseModel: db.BaseModel{
            Connection: "default",
            Table:      "user",
            FakeDelete: true,
            FakeDelKey: "status",
        },
    }
}
