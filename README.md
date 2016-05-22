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

func main() {
	t := gotable.NewTable([]gotable.Column{
		gotable.NewColumn("test", 10),
		gotable.NewColumn("test2", 8),
	})

	t.Push(23, 42)
	t.Push("hello", "world")
	t.Push(true, false)
	t.Push(123456789987654321, "this line is too long")

	t.Print()

	/* Output:
	+----------+--------+
	|   test   | test2  |
	+----------+--------+
	|        23|      42|
	|hello     |world   |
	|true      |false   |
	|1234567899|this lin|
	+----------+--------+
	*/
}
```
