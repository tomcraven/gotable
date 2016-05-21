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
	default:
		// TODO: test
		panic("unsupported format")
	}
}

func lPad(str string, width int) string {
	strLen := len(str)
	if strLen >= width {
		return str
	}

	return strings.Repeat(" ", width-strLen) + str
}

type intCell struct {
	item   int
	column Column
}

func (c intCell) Print(output Output) {
	str := strconv.Itoa(c.item)
	output.Print(lPad(str, c.column.GetWidth()))
}
