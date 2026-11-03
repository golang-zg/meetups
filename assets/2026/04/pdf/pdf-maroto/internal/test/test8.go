package test

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/v2"
)

func Test8() {

	fmt.Printf("\n---test8---\n")

	// setup core
	mrt := maroto.New()
	fmt.Printf("Core = %+v\n", mrt)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// read document to merge
	test6Pdf, err := os.ReadFile("doc/test6.pdf")
	if err != nil {
		fmt.Printf("Cannot read test document, err = %v\n", err)
		return
	}

	// merge
	err = document.Merge(test6Pdf)
	if err != nil {
		fmt.Printf("Cannot merge document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test8.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}
