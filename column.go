package gotable

// Column contains information about a column in a Table
type Column interface {
	PrintHeader(Output)
	PrintCellAt(int, Output)
	Push(interface{})
	GetWidth() int
}

type columnImpl struct {
	header Cell
	width  int
	cells  []Cell
}

// NewColumn instantiates and returns a new instance of Column
func NewColumn(name string, width int) Column {
	newColumn := &columnImpl{
		width: width,
	}
	newColumn.header = NewAlignedCell(newColumn, name, centre)
	return newColumn
}

// PrintHeader prints the column header aligned to centre
func (c *columnImpl) PrintHeader(output Output) {
	c.header.Print(output)
}

// PrintCellAt takes an ordinal and the output interface and prints the row
func (c *columnImpl) PrintCellAt(ordinal int, output Output) {
	c.cells[ordinal].Print(output)
}

// Push appends an item to the column
func (c *columnImpl) Push(x interface{}) {
	c.cells = append(c.cells, NewCell(c, x))
}

func (c *columnImpl) GetWidth() int {
	return c.width
}
