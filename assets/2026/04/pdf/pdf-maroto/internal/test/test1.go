package test

import (
	"fmt"

	"github.com/johnfercher/maroto/v2"
)

func Test1() {

	fmt.Printf("\n---test1---\n")

	// setup core
	mrt := maroto.New()
	fmt.Printf("Core = %+v\n", mrt)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test1.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)
}
