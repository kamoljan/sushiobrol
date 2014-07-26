package json

import (
	"encoding/json"
	"log"
)

type Egg struct {
	Egg     string `json:"egg"`     //0001_bbf06d39e4dac6b4cac5ee16226f6b5f7c50f071_ACA0AC_401_638
	Baby    string `json:"baby"`    //0001_6881db255b21c864c9d1e28db50dc3b71dab5b78_ACA0AC_400_637
	Infant  string `json:"infant"`  //0001_ff41e42b0134e219bc09eddda87687822460afcf_ACA0AC_200_319
	Newborn string `json:"newborn"` //0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160
}

type Result struct {
	Newborn string `json:"newborn"` //0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160
}

type Msg struct {
	Status string      `json:"status"` //"ok"
	Result interface{} `json:"data"`   //{newborn: "0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160"}
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
