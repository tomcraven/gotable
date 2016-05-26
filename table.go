package gotable

import "strings"

type asciiChar int

const (
	topLeft asciiChar = iota
	topRight
	bottomLeft
	bottomRight

	tJuncUp
	tJuncDown
	tJuncLeft
	tJuncRight

	column
	row

	plus
	space
)

var charset = map[asciiChar]string{
	topLeft:     "┌",
	topRight:    "┐",
	bottomLeft:  "└",
	bottomRight: "┘",

	tJuncUp:    "┴",
	tJuncDown:  "┬",
	tJuncLeft:  "┤",
	tJuncRight: "├",

	column: "│",
	row:    "─",
	plus:   "┼",
	space:  " ",
}

type position int

const (
	top position = iota
	middle
	bottom
	left
	right
)

var positionChar = map[position]map[position]asciiChar{
	top: map[position]asciiChar{
		left:   topLeft,
		middle: tJuncDown,
		right:  topRight,
	},
	middle: map[position]asciiChar{
		left:   tJuncRight,
		middle: plus,
		right:  tJuncLeft,
	},
	bottom: map[position]asciiChar{
		left:   bottomLeft,
		middle: tJuncUp,
		right:  bottomRight,
	},
}

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
	t.printHorizontalSeparator(output, top)
	t.printColumnHeaders(output)
	t.printHorizontalSeparator(output, middle)
}

func (t *Table) printColumnHeaders(output Output) {
	output.Print(charset[column])
	for i := range t.columns {
		output.Print(charset[space])
		t.columns[i].PrintHeader(output)
		output.Print(charset[space]).Print(charset[column])
	}
	output.Flush()
}

func (t *Table) printHorizontalSeparator(output Output, pos position) {
	leftChar := positionChar[pos][left]
	middleChar := positionChar[pos][middle]
	rightChar := positionChar[pos][right]

	output.Print(charset[leftChar])
	for i := range t.columns {
		output.Print(charset[row]).
			Print(strings.Repeat(charset[row], t.columns[i].GetWidth())).
			Print(charset[row])

		if i == (len(t.columns) - 1) {
			output.Print(charset[rightChar])
		} else {
			output.Print(charset[middleChar])
		}
	}
	output.Flush()
}

func (t *Table) printContent(output Output) {
	for i := 0; i < t.numRows; i++ {
		output.Print(charset[column])
		for j := range t.columns {
			output.Print(charset[space])
			t.columns[j].PrintCellAt(i, output)
			output.Print(charset[space]).Print(charset[column])
		}
		output.Flush()
	}
}

func (t *Table) printFooter(output Output) {
	t.printHorizontalSeparator(output, bottom)
}
