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
var optTest7 = flag.Bool("test7", false, "Test 7")
var optTest8 = flag.Bool("test8", false, "Test 8")

func main() {

	// command line
	flag.Parse()

	// select options
	if *optHelp {
		usage()
	} else if *optTest1 {
		// test 1: single page, default configuration
		test.Test1()
	} else if *optTest2 {
		// test 2: single page, configuration a4 vertical, security
		test.Test2()
	} else if *optTest3 {
		// test 3: single page, configuration a4 vertical, margins
		test.Test3()
	} else if *optTest4 {
		// test 4: single page, configuration a4 vertical, header, footer, line
		test.Test4()
	} else if *optTest5 {
		// test 5: single page, configuration a4 vertical, header, footer, line, text, qr
		test.Test5()
	} else if *optTest6 {
		// test 6: multi page, configuration a4 vertical, custom font, header, footer, text, metrics
		test.Test6()
	} else if *optTest7 {
		// test 7: multi page, configuration a4 horizontal, fonts, header, footer, text
		test.Test7()
	} else if *optTest8 {
		// test 8: merge
		test.Test8()
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %v [OPTION] [FILE]\n", os.Args[0])
	flag.PrintDefaults()
}
