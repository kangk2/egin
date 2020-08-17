package module

import (
    "github.com/daodao97/egin/pkg/db"
)

type CommonConfigEntity struct {
    Id   int64  `json:"id"`
    Name string `json:"name"`
}

var ConfigModel db.BaseModel

func init() {
    entity := CommonConfigEntity{}
    ConfigModel = db.BaseModel{
        Database: "hyperf_admin",
        Table:    "common_config",
        Entity:   entity,
    }
}
