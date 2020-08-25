package model

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
		Connection: "default",
		Table:      "common_config",
		Entity:     entity,
	}
}
