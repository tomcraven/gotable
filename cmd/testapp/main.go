package main

import "github.com/tomcraven/gotable"

type customCell struct {
}

func (c *customCell) Print(output gotable.Output) {
	output.Print("custom cell")
}

func main() {
	t := gotable.NewTable([]gotable.Column{
		gotable.NewColumn("test", 10),
		gotable.NewColumn("test2", 20),
	})

	t.Push(23, 42)
	t.Push("hello", "world")
	t.Push(true, false)
	t.Push(123456789987654321, "this line is too long")
	t.Push(23.42, 42.23)
	t.Push(&customCell{}, &customCell{})

	t.Print()

	/* Output:
	+------------+----------------------+
	|    test    |        test2         |
	+------------+----------------------+
	|         23 |                   42 |
	| hello      | world                |
	| true       | false                |
	| 1234567899 | this line is too lon |
	|      23.42 |                42.23 |
	| custom cel | custom cell          |
	+------------+----------------------+
	*/
}
