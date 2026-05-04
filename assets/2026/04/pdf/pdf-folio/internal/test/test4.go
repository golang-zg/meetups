package test

import (
	"fmt"

	"github.com/carlos7ags/folio/document"
	"github.com/carlos7ags/folio/font"
	"github.com/carlos7ags/folio/layout"
)

func Test4() {

	fmt.Printf("\n---test4---\n")

	// create document
	doc := document.NewDocument(document.PageSizeA4)
	doc.Info.Title = "Simple report"
	doc.Info.Author = "Branko"
	doc.Info.Producer = "pdf"
	doc.Info.Subject = "Meetup test"
	doc.Info.Keywords = "Golang, Zagreb, Meetup, 2026"

	// margins
	m := layout.Margins{
		Top:    25,
		Right:  25,
		Bottom: 25,
		Left:   25,
	}
	doc.SetMargins(m)

	// header
	headerFn := func(ctx document.PageContext) layout.Element {
		text := "Header"
		p := layout.NewParagraph(text, font.HelveticaBold, 12)
		p.SetSpaceAfter(10)
		return p
	}
	doc.SetHeaderElement(headerFn)

	// footer
	footerFn := func(ctx document.PageContext) layout.Element {
		text := fmt.Sprintf("Page %d/%d", ctx.PageIndex+1, ctx.TotalPages)
		p := layout.NewParagraph(text, font.Helvetica, 9)
		p.SetAlign(layout.AlignCenter)
		return p
	}
	doc.SetFooterElement(footerFn)

	// heading
	doc.Add(spacer(50))
	h1 := layout.NewHeading("Hello GolangZG!", layout.H1)
	doc.Add(h1)

	// paragraph
	p1 := layout.NewParagraph("The Go programming language is an open source project to make programmers more productive.", font.Helvetica, 12)
	doc.Add(p1)

	p2 := layout.NewParagraph("Go is expressive, concise, clean, and efficient.", font.Helvetica, 12)
	doc.Add(p2)
	doc.Add(spacer(10))

	// add rows
	for i := range 20 {
		text := fmt.Sprintf("%.5d ...1.........2.........3.........4.........5.........6.........7", i)
		p := layout.NewParagraph(text, font.Helvetica, 12)
		doc.Add(p)

		if i == 10 {
			p = layout.NewParagraph("page break", font.Helvetica, 12)
			doc.Add(p)
			// force page break
			doc.Add(layout.NewAreaBreak())
		}
	}

	// save document
	fileName := "doc/test4.pdf"
	err := doc.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)
}
