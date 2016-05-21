package gotable

import "strings"

// Column contains information about a column in a Table
type Column struct {
	name  string
	width int
	cells []cell
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
	// TODO: can this be done with cells?

	nameLength := len(c.name)
	spareRoom := c.getWidth() - nameLength
	leftPadding := strings.Repeat(" ", spareRoom/2)
	rightPadding := strings.Repeat(" ", spareRoom-(spareRoom/2))
	output.Print(leftPadding + c.name + rightPadding)
}

// PrintCellAt takes an ordinal and the output interface and prints the row
func (c *Column) PrintCellAt(ordinal int, output Output) {
	// TODO: bounds checking
	c.cells[ordinal].print(output)
}

// Push appends an item to the column
func (c *Column) Push(x interface{}) {
	c.cells = append(c.cells, newCell(c, x))
}

func (c *Column) getWidth() int {
	return c.width
}
