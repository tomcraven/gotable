package gotable

import (
	"strconv"
	"strings"
)

// Cell a single cell in a column/row
type Cell interface {
	Print(Output)
}

// New cell creates a new cell based off the input type
func NewCell(column Column, x interface{}) Cell {
	switch x.(type) {
	case int:
		return intCell{
			item:   x.(int),
			column: column,
		}
	case string:
		return stringCell{
			item:   x.(string),
			column: column,
		}
	default:
		panic("unsupported format")
	}
}

// --------------------
// Helpers
// --------------------

func lPad(str string, width int) string {
	strLen := len(str)
	if strLen >= width {
		return str
	}

	return strings.Repeat(" ", width-strLen) + str
}

func rPad(str string, width int) string {
	strLen := len(str)
	if strLen >= width {
		return str
	}

	return str + strings.Repeat(" ", width-strLen)
}

// --------------------
// Cells
// --------------------

type intCell struct {
	item   int
	column Column
}

func (c intCell) Print(output Output) {
	str := strconv.Itoa(c.item)
	output.Print(lPad(str, c.column.GetWidth()))
}

// --------------------

type stringCell struct {
	item   string
	column Column
}

func (c stringCell) Print(output Output) {
	output.Print(rPad(c.item, c.column.GetWidth()))
}
