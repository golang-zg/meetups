package test

import (
	"fmt"

	"github.com/carlos7ags/folio/document"
	"github.com/carlos7ags/folio/font"
	"github.com/carlos7ags/folio/layout"
)

func Test1() {

	fmt.Printf("\n---test1---\n")

	// create document
	doc := document.NewDocument(document.PageSizeA4)

	// heading
	h1 := layout.NewHeading("Hello GolangZG!", layout.H1)
	doc.Add(h1)

	// paragraph
	p1 := layout.NewParagraph("The Go programming language is an open source project to make programmers more productive.", font.Helvetica, 12)
	doc.Add(p1)

	p2 := layout.NewParagraph("Go is expressive, concise, clean, and efficient.", font.Helvetica, 12)
	doc.Add(p2)

	// save document
	fileName := "doc/test1.pdf"
	err := doc.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)
}
