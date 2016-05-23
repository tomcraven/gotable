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
	case float32:
		return createFloatCell(float64(x.(float32)), column, align)
	case float64:
		return createFloatCell(x.(float64), column, align)
	default:
		panic("unsupported cell format")
	}
}

func createFloatCell(x float64, column Column, alignment Alignment) Cell {
	return floatCell{
		baseCell: createBaseCell(column, alignment, Right),
		item:     x,
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
	outputAligned := NewOutputAligned(c.column.GetWidth(), output, c.alignment)
	outputAligned.Print(str)
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

// --------------------

type floatCell struct {
	baseCell
	item float64
}

func (c floatCell) Print(output Output) {
	c.printString(strconv.FormatFloat(c.item, 'f', -1, 64), output)
}
