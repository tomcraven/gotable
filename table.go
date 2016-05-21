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
	numRows int
}

// NewTable creates a new table
func NewTable(columns []Column) Table {
	return Table{
		columns: columns,
	}
}

// Push appends values to the table
func (t *Table) Push(values ...interface{}) {
	for i, value := range values {
		// TODO: bounds checking
		t.columns[i].Push(value)
	}
	t.numRows++
}

// Print outputs the table to the output object
func (t *Table) Print(output Output) {
	outputBuffered := NewOutputBuffered(output)
	t.printHeader(&outputBuffered)

	for i := 0; i < t.numRows; i++ {
		outputBuffered.Print(columnChar)
		for _, column := range t.columns {
			column.PrintCellAt(i, &outputBuffered)
			outputBuffered.Print(columnChar)
		}
		outputBuffered.Flush()
	}
	t.printHorizontalSeparator(&outputBuffered)
}

func (t *Table) printHeader(output Output) {
	t.printHorizontalSeparator(output)
	t.printColumnHeaders(output)
	t.printHorizontalSeparator(output)
}

func (t *Table) printColumnHeaders(output Output) {
	output.Print(columnChar)
	for _, column := range t.columns {
		column.PrintHeader(output)
		output.Print(columnChar)
	}
	output.Flush()
}

func (t *Table) printHorizontalSeparator(output Output) {
	output.Print(cornerChar)
	for _, column := range t.columns {
		output.Print(strings.Repeat(rowChar, column.getWidth())).
			Print(cornerChar)
	}
	output.Flush()
}
