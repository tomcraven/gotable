package gotable

import "fmt"

// Output interface for injecting different output formats
type Output interface {
	Print(string)
}

// OutputStdOut prints all messages to std out
type OutputStdOut struct {
}

// Print - prints to std out
func (output *OutputStdOut) Print(a string) {
	fmt.Println(a)
}

// NullOutput - null implementation
type NullOutput struct {
}

// Print - null implementation
func (n *NullOutput) Print(a string) {
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
func (o *OutputBuffered) Print(a string) {
	o.cumulativeOutput += a
}

// Flush sends a concatenated version of the previous calls to Print since
// object creation or the last call to Flush
func (o *OutputBuffered) Flush() {
	o.otherOutput.Print(o.cumulativeOutput)
	o.cumulativeOutput = ""
}
