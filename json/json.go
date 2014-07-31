package json

import (
	"encoding/json"
	"log"
)

type Msg struct {
	Status string      `json:"status"` //"ok"
	Result interface{} `json:"data"`   //{newborn: "0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160"}
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

type Msg3 struct {
	Status  string      `json:"status"`  //"OK" || "ERROR"
	Result  interface{} `json:"data"`    //{newborn: "0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160"}
	Message string      `json:"message"` //"Some Error is happen bla bla"
}

func Message3(status string, result interface{}, message string) []byte {
	m := Msg3{
		Status:  status,
		Result:  result,
		Message: message,
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("Unable to json.Marshal ", err)
	}
	return b
}
