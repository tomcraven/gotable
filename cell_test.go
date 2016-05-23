package gotable_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/tomcraven/gotable"
	. "github.com/tomcraven/gotable/gotable_mock"
)

var _ = Describe("Cell", func() {
	Describe("NewCell", func() {
		It("panics with unsupported cell formats", func() {
			type randomType struct{}
			Expect(func() {
				NewCell(NewColumn("1", 1), randomType{})
			}).Should(Panic())
		})
	})

	Describe("Print", func() {
		var (
			c              Cell
			mockController *gomock.Controller
			mockColumn     *MockColumn
			mockOutput     *MockOutput
		)

		BeforeEach(func() {
			mockController = gomock.NewController(GinkgoT())
			mockOutput = NewMockOutput(mockController)
			mockColumn = NewMockColumn(mockController)
		})

		type printConfiguration struct {
			input          interface{}
			width          int
			expectedOutput string
		}

		printTest := func(input interface{}, width int, expectedOutput string) {
			c = NewCell(mockColumn, input)
			mockColumn.EXPECT().GetWidth().AnyTimes().Return(width)
			mockOutput.EXPECT().Print(expectedOutput)
			c.Print(mockOutput)
		}

		DescribeTable("intCell", printTest,
			Entry("padded left", 1, 4, "   1"),
			Entry("not padded", 1234, 4, "1234"),
			Entry("truncated", 12345, 4, "1234"),
		)

		DescribeTable("stringCell", printTest,
			Entry("padded right", "hello", 10, "hello     "),
			Entry("no padding", "world", 5, "world"),
			Entry("truncated", "blah", 2, "bl"),
		)

		DescribeTable("boolCell", printTest,
			Entry("padding right true", true, 6, "true  "),
			Entry("no padding true", true, 4, "true"),
			Entry("padding right false", false, 10, "false     "),
			Entry("no padding false", false, 5, "false"),
			Entry("truncated", true, 3, "tru"),
		)

		// Floats are tricky to test, a float32(1.1) looks like this on my machine:
		//   1.100000023841858
		// Where as a float64(1.1) looks like this:
		//   1.1
		// For now, just test that printing with no padding is working
		DescribeTable("floatCell", printTest,
			Entry("padding left float32", float32(1.1), 3, "1.1"),
			Entry("padding left float64", float64(1.1), 3, "1.1"),
		)
	})
})
