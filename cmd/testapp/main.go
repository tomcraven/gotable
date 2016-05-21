package main

import "github.com/tomcraven/gotable"

func main() {
	t := gotable.NewTable([]gotable.Column{
		gotable.NewColumn("test", 10),
		gotable.NewColumn("test2", 8),
	})

	t.Push(1, 2)
	t.Push(123, 45342)

	output := gotable.OutputStdOut{}
	t.Print(&output)
}
