package main

import "github.com/tomcraven/gotable"

type customCell struct {
}

func (c *customCell) Print(output gotable.Output) {
	output.Print("custom cell")
}

func main() {
	t := gotable.NewTable([]gotable.Column{
		gotable.NewColumn("column 1", 11),
		gotable.NewColumn("column 2", 20),
	})

	t.Push(23, 42)
	t.Push("hello", "world")
	t.Push(true, false)
	t.Push(123456789987654321, "this line is too long")
	t.Push(23.42, 42.23)
	t.Push(&customCell{}, &customCell{})
	t.Push("blank right", nil)
	t.Push(nil, "blank left")

	t.Print()

	/* Output:
	┌─────────────┬──────────────────────┐
	│  column 1   │       column 2       │
	├─────────────┼──────────────────────┤
	│          23 │                   42 │
	│ hello       │ world                │
	│ true        │ false                │
	│ 12345678998 │ this line is too lon │
	│       23.42 │                42.23 │
	│ custom cell │ custom cell          │
	│ blank right │                      │
	│             │ blank left           │
	└─────────────┴──────────────────────┘
	*/
}
