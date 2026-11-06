package http

import (
	"encoding/json"
	"time"
)

type TaskDTO struct {
	title       string
	description string
}

type ErrorDTO struct {
	Massage string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
}
