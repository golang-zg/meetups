// pdf - PDF
//
// Remarks
//   - none
//
// See pdf.txt for program notes.
//
// See CHANGELOG.txt for revision history.
package main

import (
	"flag"
	"fmt"
	"os"
	"pdf/internal/test"
)

// imports

// typedefs

// common declarations

var optHelp = flag.Bool("help", false, "Show help")

// option, flag true/false
var optTest1 = flag.Bool("test1", false, "Test 1")
var optTest2 = flag.Bool("test2", false, "Test 2")
var optTest3 = flag.Bool("test3", false, "Test 3")
var optTest4 = flag.Bool("test4", false, "Test 4")
var optTest5 = flag.Bool("test5", false, "Test 5")
var optTest6 = flag.Bool("test6", false, "Test 6")

func main() {

	// command line
	flag.Parse()

	// select options
	if *optHelp {
		usage()
	} else if *optTest1 {
		// test 1: single page, portrait
		test.Test1()
	} else if *optTest2 {
		// test 2: single page, portrait, security
		test.Test2()
	} else if *optTest3 {
		// test 3: single page, portrait, margins
		test.Test3()
	} else if *optTest4 {
		// test 4: multi page, portrait, margins, header, footer
		test.Test4()
	} else if *optTest5 {
		// test 5: single page, portrait, margins, header, footer, bookmarks
		test.Test5()
	} else if *optTest6 {
		// test 6: single page, portrait, margins, header, footer, html
		test.Test6()
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %v [OPTION] [FILE]\n", os.Args[0])
	flag.PrintDefaults()
}
