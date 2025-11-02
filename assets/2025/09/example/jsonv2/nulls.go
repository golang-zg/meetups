package main

import (
	"encoding/json/v2"
	"fmt"
)

type Data struct {
	List []int          `json:"list"`
	Map  map[string]int `json:"map"`
}

func modeNullsNew() {
	var d Data
	nullsNew(d)
}

func modeNullsOld() {
	var d Data
	nullsOld(d)
}

func nullsNew(d Data) {
	b, err := json.Marshal(d)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}

func nullsOld(d Data) {
	b, err := json.Marshal(d,
		json.FormatNilSliceAsNull(true),
		json.FormatNilMapAsNull(true),
	)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}
