package gotable

import "strings"

const (
	columnChar      = "|"
	rowChar         = "-"
	cornerChar      = "+"
	columnSeparator = " "
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
func (t *Table) Print() {
	t.PrintTo(&OutputStdOut{})
}

// PrintTo takes an instance of the output interface and prints to it
func (t *Table) PrintTo(output Output) {
	outputBuffered := NewOutputBuffered(output)
	t.printHeader(&outputBuffered)
	t.printContent(&outputBuffered)
	t.printFooter(&outputBuffered)
}

func (t *Table) printHeader(output Output) {
	t.printHorizontalSeparator(output)
	t.printColumnHeaders(output)
	t.printHorizontalSeparator(output)
}

func (t *Table) printColumnHeaders(output Output) {
	output.Print(columnChar)
	for i := range t.columns {
		output.Print(columnSeparator)
		t.columns[i].PrintHeader(output)
		output.Print(columnSeparator).Print(columnChar)
	}
	output.Flush()
}

func (t *Table) printHorizontalSeparator(output Output) {
	output.Print(cornerChar)
	for i := range t.columns {
		output.Print(rowChar).
			Print(strings.Repeat(rowChar, t.columns[i].GetWidth())).
			Print(rowChar).
			Print(cornerChar)
	}
	output.Flush()
}

func (t *Table) printContent(output Output) {
	for i := 0; i < t.numRows; i++ {
		output.Print(columnChar)
		for j := range t.columns {
			output.Print(columnSeparator)
			t.columns[j].PrintCellAt(i, output)
			output.Print(columnSeparator).Print(columnChar)
		}
		output.Flush()
	}
}

func (t *Table) printFooter(output Output) {
	t.printHorizontalSeparator(output)
}
