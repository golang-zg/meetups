package main

import (
	"encoding/json/v2"
	"fmt"
)

type Image struct {
	Data []byte `json:"data"`
}

type ImageBase64 struct {
	Data []byte `json:"data,format:base64"`
}

type ImageHex struct {
	Data []byte `json:"data,format:hex"`
}

type ImageArray struct {
	Data []byte `json:"data,format:array"`
}

func modeByte(mode string) {
	img := Image{
		Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
	}
	switch mode {
	case "byte-notag":
		image(img)
	case "byte-base64":
		imageBase64(img)
	case "byte-hex":
		imageHex(img)
	case "byte-array":
		imageArray(img)
	default:
		panic("unknown mode")
	}
}

func image(img Image) {
	b, err := json.Marshal(img)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}

func imageBase64(img Image) {
	b, err := json.Marshal(ImageBase64{
		Data: img.Data,
	})
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}

func imageHex(img Image) {
	b, err := json.Marshal(ImageHex{
		Data: img.Data,
	})
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}

func imageArray(img Image) {
	b, err := json.Marshal(ImageArray{
		Data: img.Data,
	})
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
}
