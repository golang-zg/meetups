package test

import (
	"fmt"

	"github.com/carlos7ags/folio/document"
	"github.com/carlos7ags/folio/font"
	"github.com/carlos7ags/folio/layout"
)

func Test2() {

	fmt.Printf("\n---test2---\n")

	// create document
	doc := document.NewDocument(document.PageSizeA4)
	doc.Info.Title = "Simple report"
	doc.Info.Author = "Branko"
	doc.Info.Producer = "pdf"
	doc.Info.Subject = "Meetup test"
	doc.Info.Keywords = "Golang, Zagreb, Meetup, 2026"

	// set password
	enc := document.EncryptionConfig{
		UserPassword:  "1234",
		OwnerPassword: "abcd",
	}
	doc.SetEncryption(enc)

	// heading
	h1 := layout.NewHeading("Hello GolangZG!", layout.H1)
	doc.Add(h1)

	// paragraph
	p1 := layout.NewParagraph("The Go programming language is an open source project to make programmers more productive.", font.Helvetica, 12)
	doc.Add(p1)

	p2 := layout.NewParagraph("Go is expressive, concise, clean, and efficient.", font.Helvetica, 12)
	doc.Add(p2)

	// save document
	fileName := "doc/test2.pdf"
	err := doc.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)
}
