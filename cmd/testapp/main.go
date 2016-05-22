package main

import "github.com/tomcraven/gotable"

func main() {
	t := gotable.NewTable([]gotable.Column{
		gotable.NewColumn("test", 10),
		gotable.NewColumn("test2", 8),
	})

	t.Push(23, 42)
	t.Push("hello", "world")
	t.Push(true, false)

	output := gotable.OutputStdOut{}
	t.Print(&output)

	/* Outputs:
	+----------+--------+
	|   test   | test2  |
	+----------+--------+
	|        23|      42|
	|hello     |world   |
	|true      |false   |
	+----------+--------+
	*/
}
