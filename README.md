[![Build Status](https://travis-ci.org/tomcraven/gotable.svg?branch=master)](https://travis-ci.org/tomcraven/gotable)
[![Coverage Status](https://coveralls.io/repos/tomcraven/gotable/badge.svg?branch=master&service=github)](https://coveralls.io/github/tomcraven/gotable?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomcraven/gotable)](https://goreportcard.com/report/github.com/tomcraven/gotable)
# gotable

A simple ASCII table renderer for fun (and so i can play around with ginkgo...)

### Installation

Library alone:

```
go get github.com/tomcaven/gotable
```

Running the tests:

```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get github.com/golang/mock/gomock
```

### Example

```
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
```
