package gotable

import "strings"

// Column contains information about a column in a Table
type Column interface {
	PrintHeader(Output)
	PrintCellAt(int, Output)
	Push(interface{})
	getWidth() int
}

type columnImpl struct {
	name  string
	width int
	cells []cell
}

// NewColumn instantiates and returns a new instance of Column
func NewColumn(name string, width int) Column {
	return &columnImpl{
		name:  name,
		width: width,
	}
}

// PrintHeader prints the column header
// Currently only supports aligning to centre
func (c *columnImpl) PrintHeader(output Output) {
	// TODO: can this be done with cells?

	nameLength := len(c.name)
	spareRoom := c.getWidth() - nameLength
	leftPadding := strings.Repeat(" ", spareRoom/2)
	rightPadding := strings.Repeat(" ", spareRoom-(spareRoom/2))
	output.Print(leftPadding + c.name + rightPadding)
}

// PrintCellAt takes an ordinal and the output interface and prints the row
func (c *columnImpl) PrintCellAt(ordinal int, output Output) {
	// TODO: bounds checking
	c.cells[ordinal].print(output)
}

// Push appends an item to the column
func (c *columnImpl) Push(x interface{}) {
	c.cells = append(c.cells, newCell(c, x))
}

func (c *columnImpl) getWidth() int {
	return c.width
}
