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
