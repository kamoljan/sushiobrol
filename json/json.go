package json

import (
	"encoding/json"
	"log"
)

type Msg struct {
	Status string      `json:"status"`
	Result interface{} `json:"data"`
}

type Result struct {
	Image []Fid `json:"image"`
}

type Fid struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func Message(status string, result interface{}) []byte {
	m := Msg{
		Status: status,
		Result: result,
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("Unable to json.Marshal ", err)
	}
	return b
}
