package entity

import (
	"github.com/google/uuid"
)

type Author struct {
	Id   string `json:"id"`
	Name string `json: "name"`
}

func (a *Author) SetName(name string) {
	a.Id = uuid.NewString()
	a.Name = name
}
