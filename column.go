package gotable

import "strings"

// Column contains information about a column in a Table
type Column struct {
	name  string
	width int
}

// NewColumn instantiates and returns a new instance of Column
func NewColumn(name string, width int) Column {
	return Column{
		name:  name,
		width: width,
	}
}

// PrintHeader prints the column header
// Currently only supports aligning to centre
func (c *Column) PrintHeader(output Output) {
	nameLength := len(c.name)
	spareRoom := c.getWidth() - nameLength
	leftPadding := strings.Repeat(" ", spareRoom/2)
	rightPadding := strings.Repeat(" ", spareRoom-(spareRoom/2))
	output.Print(leftPadding + c.name + rightPadding)
}

func (c *Column) getWidth() int {
	return c.width
}
