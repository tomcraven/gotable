package gotable

import "strings"

const (
	columnChar = "|"
	rowChar    = "-"
	cornerChar = "+"
)

// Table holds the table state
type Table struct {
	columns []Column
}

// NewTable creates a new table
func NewTable(columns []Column) Table {
	return Table{
		columns: columns,
	}
}

// Print outputs the table to the output object
func (t *Table) Print(output Output) {
	outputBuffered := NewOutputBuffered(output)
	t.printHeader(outputBuffered)
}

func (t *Table) printHeader(output OutputBuffered) {
	t.printHorizontalSeparator(output)
	t.printColumnHeaders(output)
	t.printHorizontalSeparator(output)
}

func (t *Table) printColumnHeaders(output OutputBuffered) {
	output.Print(columnChar)
	for _, column := range t.columns {
		column.PrintHeader(&output)
		output.Print(columnChar)
	}
	output.Flush()
}

func (t *Table) printHorizontalSeparator(output OutputBuffered) {
	output.Print(cornerChar)
	for _, column := range t.columns {
		output.Print(strings.Repeat(rowChar, column.getWidth())).
			Print(cornerChar)
	}
	output.Flush()
}
