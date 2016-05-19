package gotable_test

import (
	. "github.com/tomcraven/gotable"
)

var _ = Describe("Table", func() {
	It("creates a blank table", func() {
		_ = NewTable([]Column{})
	})
})
