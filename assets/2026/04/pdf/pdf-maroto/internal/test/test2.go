package test

import (
	"fmt"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
)

func Test2() {

	fmt.Printf("\n---test2---\n")

	// setup configuration
	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithAuthor("Branko", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	builder.WithDebug(true) // debug = true creates frames
	builder.WithProtection(protection.Copy, "root", "1234")
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	fmt.Printf("Core = %+v\n", mrt)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test2.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}
