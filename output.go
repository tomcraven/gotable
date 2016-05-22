package gotable

import "fmt"

// Output interface for injecting different output formats
type Output interface {
	Print(string) Output
	Flush() Output
}

// OutputStdOut prints all messages to std out
type OutputStdOut struct {
}

// Print - prints to std out
func (o *OutputStdOut) Print(a string) Output {
	fmt.Println(a)
	return o
}

// Flush - noop
func (o *OutputStdOut) Flush() Output {
	return o
}

// NullOutput - null implementation
type NullOutput struct {
}

// Print - null implementation
func (n *NullOutput) Print(a string) Output {
	return n
}

// Flush - null implementation
func (n *NullOutput) Flush() Output {
	return n
}

// OutputBuffered - an output implementation that takes another instance
// out Output and buffers up the calls to Print until flush is called
type OutputBuffered struct {
	cumulativeOutput string
	otherOutput      Output
}

// NewOutputBuffered creates a new OutputBuffered
func NewOutputBuffered(otherOutput Output) OutputBuffered {
	return OutputBuffered{
		cumulativeOutput: "",
		otherOutput:      otherOutput,
	}
}

// Print buffers the output until Flush is called
func (o *OutputBuffered) Print(a string) Output {
	o.cumulativeOutput += a
	return o
}

// Flush sends a concatenated version of the previous calls to Print since
// object creation or the last call to Flush
func (o *OutputBuffered) Flush() Output {
	o.otherOutput.Print(o.cumulativeOutput)
	o.cumulativeOutput = ""
	return o
}

// OutputAligned takes an output object and guarantees that when flush is called
// the data sent to the consumer will be x characters long (by padding or
// truncating)
type OutputAligned struct {
	otherOutput Output
	fixedWidth  int
	alignment   Alignment
}

// NewOutputAligned creates a new OutputAligned
func NewOutputAligned(fixedWidth int, otherOutput Output, alignment Alignment) OutputAligned {
	return OutputAligned{
		otherOutput: otherOutput,
		fixedWidth:  fixedWidth,
		alignment:   alignment,
	}
}

// Print guarantees that the string has the same length as the width the
// OutputAligned has been configured with
func (o *OutputAligned) Print(str string) Output {
	strlen := len(str)
	if strlen <= o.fixedWidth {
		o.otherOutput.Print(padForAlignment(str, o.fixedWidth, o.alignment))
	} else {
		o.otherOutput.Print(str[0:o.fixedWidth])
	}
	return o
}

// Flush is a noop in OutputAligned
func (o *OutputAligned) Flush() Output {
	return o
}
