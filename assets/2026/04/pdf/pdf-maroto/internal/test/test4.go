package test

import (
	"fmt"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func Test4() {

	fmt.Printf("\n---test4---\n")

	// setup configuration
	pn := props.PageNumber{
		Pattern: "Page {current} of {total}",
		Place:   props.RightBottom,
	}

	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithPageNumber(pn)
	builder.WithAuthor("Branko", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	builder.WithDebug(true) // debug = true creates frames
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	fmt.Printf("Core = %+v\n", mrt)

	// header
	headerRow1 := text.NewRow(5, "Sales")
	headerRow2 := text.NewRow(5, "Simple report", props.Text{
		Style: fontstyle.Bold,
		Align: align.Center,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})
	headerRow3 := text.NewRow(5, "Simple report", props.Text{
		Style: fontstyle.Bold,
		Align: align.Right,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})

	err := mrt.RegisterHeader(headerRow1, headerRow2, headerRow3)
	if err != nil {
		fmt.Printf("Cannot register document header, err = %v\n", err)
		return
	}

	// footer
	footerRow1 := text.NewRow(5, "(c) POINTER d.o.o.", props.Text{
		Size:  5,
		Align: align.Left,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})

	err = mrt.RegisterFooter(footerRow1)
	if err != nil {
		fmt.Printf("Cannot register document footer, err = %v\n", err)
		return
	}

	// line
	mrt.AddRow(5, line.NewCol(12))
	mrt.AddRow(15, line.NewCol(11))
	mrt.AddRow(5, line.NewCol(12, props.Line{Style: linestyle.Dashed, Color: &props.BlueColor, Thickness: 1}))

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test4.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}
