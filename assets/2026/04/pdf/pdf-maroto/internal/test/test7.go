package test

import (
	"fmt"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

func Test7() {

	fmt.Printf("\n---test7---\n")

	const LINE_SIZE = 4

	// setup fonts
	fontFamily := "courier-new"
	fontRepo := repository.New()
	// fontRepo.AddUTF8Font(fontFamily, fontstyle.Normal, "./fonts/RobotoMono-Regular.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Normal, "./fonts/RobotoMono-Light.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Italic, "./fonts/RobotoMono-Italic.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Bold, "./fonts/RobotoMono-Bold.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.BoldItalic, "./fonts/RobotoMono-BoldItalic.ttf")

	fonts, err := fontRepo.Load()
	if err != nil {
		fmt.Printf("Cannot load fonts, err = %v\n", err)
		return
	}
	fmt.Printf("Fonts = %+v\n", fonts)

	// setup configuration
	pn := props.PageNumber{
		Pattern: "APP 1.2.3 | {current}/{total}",
		Place:   props.RightBottom,
	}

	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Horizontal)
	builder.WithPageNumber(pn)
	builder.WithAuthor("Branko", true)
	builder.WithCreator("APP 1.2.3", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	// builder.WithDebug(true) // debug = true creates frames
	builder.WithCustomFonts(fonts)
	builder.WithDefaultFont(&props.Font{Family: fontFamily})
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	mrt = maroto.NewMetricsDecorator(mrt)
	fmt.Printf("Core = %+v\n", mrt)

	// header
	// .........1.........2.........3.........4.........5.........6.........7.........8.........9.........0.........1.........2.........3
	// 1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
	// xxxxxxxxx1: xxxxxxxxx1xxxxxxxxx2 							         										  yyyy-mm-dd hh:mm:ss
	//
	headerRow1 := row.New(LINE_SIZE)
	headerRow1 = headerRow1.Add(
		text.NewCol(6, fmt.Sprintf("%s: %s ", "Sales", "Simple report")),
		text.NewCol(6, time.Now().Format("2006-01-02 15:04:05"), props.Text{
			Align: align.Right,
		}))
	headerRowDelimiter := text.NewRow(LINE_SIZE, "")

	err = mrt.RegisterHeader(headerRow1, headerRowDelimiter, headerRowDelimiter)
	if err != nil {
		fmt.Printf("Cannot register document header, err = %v\n", err)
		return
	}

	// footer
	footerRowDelimiter := text.NewRow(LINE_SIZE, "")
	err = mrt.RegisterFooter(footerRowDelimiter)
	if err != nil {
		fmt.Printf("Cannot register document footer, err = %v\n", err)
		return
	}

	// template row
	mrt.AddRow(LINE_SIZE, text.NewCol(12, ".........1.........2.........3.........4.........5.........6.........7.........8.........9.........0.........1.........2.........3"))
	mrt.AddRow(LINE_SIZE, text.NewCol(12, "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"))

	// add rows
	for i := 0; i < 500; i++ {
		mrt.AddRow(LINE_SIZE, text.NewCol(12, fmt.Sprintf("%5d", i)))
	}

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test7.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}
