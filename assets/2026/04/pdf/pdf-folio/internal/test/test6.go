package test

import (
	"fmt"

	"github.com/carlos7ags/folio/document"
	"github.com/carlos7ags/folio/font"
	"github.com/carlos7ags/folio/html"
	"github.com/carlos7ags/folio/layout"
)

func Test6() {

	fmt.Printf("\n---test6---\n")

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

	// html
	// Define your HTML content
	htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body { font-family: Helvetica; }
				table { border-collapse: collapse; width: 100%; }
				th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
			</style>
		</head>
		<body>
			<h1>Invoice</h1>
			<h2>Invoice number: 2026-123456</h2>
			<p>Bill to: <strong>GolangZG</strong></p>
			<table>
				<tr>
					<th>Item</th>
					<th>Amount</th>
				</tr>
				<tr>
					<td>Consulting</td>
					<td>1.200,00 EUR</td>
				</tr>
				<tr>
					<td>Development</td>
					<td>2.500,00 EUR</td>
				</tr>
			</table>
			<p>Thank you for your order!</p>
		</body>
		</html>
	`

	// convert html
	elements, err := html.Convert(htmlContent, nil)
	if err != nil {
		fmt.Printf("Cannot convert html, err = %v\n", err)
		return
	}

	// add elements to document
	for _, elem := range elements {
		doc.Add(elem)
	}

	fileName := "doc/test6.pdf"
	err = doc.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)
}
