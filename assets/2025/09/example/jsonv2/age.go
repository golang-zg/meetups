package main

import (
	"encoding/json/v2"
	"fmt"
	"time"
)

type Age struct {
	DOB time.Time `json:"dob,format:DateOnly"`
}

func modeAge() {
	a := Age{
		DOB: time.Date(
			2025, 9, 23, 18, 35, 0, 0, time.UTC),
	}
	b, err := json.Marshal(a)
	fmt.Println(string(b))

	if err != nil {
		fmt.Println(err)
	}
}
