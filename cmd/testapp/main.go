package main

import "github.com/tomcraven/gotable"

func main() {
	t := gotable.NewTable([]gotable.Column{
		gotable.NewColumn("test", 10),
		gotable.NewColumn("test2", 8),
	})

	t.Push(1, 2)
	t.Push("abc", 123)
	t.Push("hello", "world")

	output := gotable.OutputStdOut{}
	t.Print(&output)

	/* Outputs:
	+----------+--------+
	|   test   | test2  |
	+----------+--------+
	|         1|       2|
	|abc       |     123|
	|hello     |world   |
	+----------+--------+
	*/
}
