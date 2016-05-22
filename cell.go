package gotable

import "strconv"

// Cell a single cell in a column/row
type Cell interface {
	Print(Output)
}

// NewCell creates a new cell based off the input type
func NewCell(column Column, x interface{}) Cell {
	return NewAlignedCell(column, x, unset)
}

// NewAlignedCell creates a cell with a specified alignment
// pass 'unset' to use default alignment, or simply call NewCell
func NewAlignedCell(column Column, x interface{}, align Alignment) Cell {
	switch x.(type) {
	case int:
		return intCell{
			baseCell: createBaseCell(column, align, Right),
			item:     x.(int),
		}
	case string:
		return stringCell{
			baseCell: createBaseCell(column, align, Left),
			item:     x.(string),
		}
	case bool:
		return boolCell{
			baseCell: createBaseCell(column, align, Left),
			item:     x.(bool),
		}
	default:
		panic("unsupported cell format")
	}
}

func createBaseCell(c Column, alignment, defaultAlignment Alignment) baseCell {
	return baseCell{
		column:    c,
		alignment: getAlignment(alignment, defaultAlignment),
	}
}

func getAlignment(a, backup Alignment) Alignment {
	if a == unset {
		return backup
	}
	return a
}

// --------------------
// Cells
// --------------------

type baseCell struct {
	column    Column
	alignment Alignment
}

func (c *baseCell) printString(str string, output Output) {
	output.Print(
		padForAlignment(str, c.column.GetWidth(), c.alignment),
	)
}

// --------------------

type intCell struct {
	baseCell
	item int
}

func (c intCell) Print(output Output) {
	str := strconv.Itoa(c.item)
	c.printString(str, output)
}

// --------------------

type stringCell struct {
	baseCell
	item string
}

func (c stringCell) Print(output Output) {
	c.printString(c.item, output)
}

// --------------------

type boolCell struct {
	baseCell
	item bool
}

func (c boolCell) Print(output Output) {
	if c.item {
		c.printString("true", output)
	} else {
		c.printString("false", output)
	}
}
