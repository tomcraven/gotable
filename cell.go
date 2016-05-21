package gotable

import (
	"strconv"
	"strings"
)

type cell interface {
	print(Output)
}

func newCell(column Column, x interface{}) cell {
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

func (c intCell) print(output Output) {
	str := strconv.Itoa(c.item)
	output.Print(lPad(str, c.column.getWidth()))
}
