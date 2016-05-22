package gotable

import (
	"strconv"
	"strings"
)

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

func padForAlignment(str string, width int, a Alignment) string {
	switch a {
	case Left:
		return rPad(str, width)
	case Right:
		return lPad(str, width)
	case Centre:
		return centrePad(str, width)
	default:
		panic("unsupported alignment " + strconv.Itoa(int(a)))
	}
}

type Alignment int

const (
	// Left align
	Left Alignment = iota
	// Right align
	Right
	// Centre align
	Centre
	unset
)
