package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"os"
)

func main() {
	mode := ""
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	person := Person{
		Name:             "Zlatko",
		AttendingMeetups: true,
	}

	switch mode {
	case "age":
		modeAge()
	case "null-new":
		modeNullsNew()
	case "null-old":
		modeNullsOld()
	case "multiline":
		multiline(person)
	case "byte-notag", "byte-base64", "byte-hex", "byte-array":
		modeByte(mode)
	default:
		panic("unknown mode")
	}
}

func new(person Person) {
	b, err := json.Marshal(person)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}

func old(person Person) {
	b, err := json.Marshal(person,
		json.FormatNilSliceAsNull(true),
		json.FormatNilMapAsNull(true),
	)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}

func multiline(person Person) {
	b, err := json.Marshal(person,
		jsontext.Multiline(true),
	)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}
