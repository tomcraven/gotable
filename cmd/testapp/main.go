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
	t.Push(123456789987654321, "this line is too long")
	t.Push(23.42, 42.23)

	t.Print()

	/* Output:
	+----------+--------+
	|   test   | test2  |
	+----------+--------+
	|        23|      42|
	|hello     |world   |
	|true      |false   |
	|1234567899|this lin|
	|     23.42|   42.23|
	+----------+--------+
	*/
}
