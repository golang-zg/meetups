// test - Test API
//
// Remarks
//   - none
//
// See CHANGELOG.txt for revision history.
package test

import (
	"github.com/carlos7ags/folio/font"
	"github.com/carlos7ags/folio/layout"
)

// imports

// typedefs

// common declarations

func spacer(points float64) *layout.Paragraph {
	p := layout.NewParagraph(" ", font.Helvetica, 1)
	p.SetSpaceBefore(points)
	return p
}
