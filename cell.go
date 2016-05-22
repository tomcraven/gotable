package gotable

import (
	"strconv"
	"strings"
)

// Cell a single cell in a column/row
type Cell interface {
	Print(Output)
	SetAlignment(alignment)
}

// NewCell creates a new cell based off the input type
func NewCell(column Column, x interface{}) Cell {
	return NewAlignedCell(column, x, unset)
}

// NewAlignedCell creates a cell with a specified alignment
// pass 'unset' to use default alignment, or simply call NewCell
func NewAlignedCell(column Column, x interface{}, align alignment) Cell {
	switch x.(type) {
	case int:
		return intCell{
			baseCell: baseCell{
				column:    column,
				alignment: getAlignment(align, right),
			},
			item: x.(int),
		}
	case string:
		return stringCell{
			baseCell: baseCell{
				column:    column,
				alignment: getAlignment(align, left),
			},
			item: x.(string),
		}
	default:
		panic("unsupported cell format")
	}
}

func getAlignment(a, backup alignment) alignment {
	if a == unset {
		return backup
	}
	return a
}

// --------------------
// Cell helpers
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

func centrePad(str string, width int) string {
	strLen := len(str)
	if strLen >= width {
		return str
	}

	spareRoom := width - strLen
	leftPadding := strings.Repeat(" ", spareRoom/2)
	rightPadding := strings.Repeat(" ", spareRoom-(spareRoom/2))
	return leftPadding + str + rightPadding
}

func padForAlignment(str string, width int, a alignment) string {
	switch a {
	case left:
		return rPad(str, width)
	case right:
		return lPad(str, width)
	case centre:
		return centrePad(str, width)
	}

	// Should not reach here but add a default return value anyway
	return centrePad(str, width)
}

type alignment int

const (
	unset alignment = iota
	left
	right
	centre
)

// --------------------
// Cells
// --------------------

type baseCell struct {
	column    Column
	alignment alignment
}

func (c baseCell) SetAlignment(a alignment) {
	c.alignment = a
}

// --------------------

type intCell struct {
	baseCell
	item int
}

func (c intCell) Print(output Output) {
	str := strconv.Itoa(c.item)
	output.Print(
		padForAlignment(str, c.column.GetWidth(), c.alignment),
	)
}

// --------------------

type stringCell struct {
	baseCell
	item string
}

func (c stringCell) Print(output Output) {
	output.Print(
		padForAlignment(c.item, c.column.GetWidth(), c.alignment),
	)
}
