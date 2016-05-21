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
			mockColumn.EXPECT().GetWidth().Return(width)
			mockOutput.EXPECT().Print(expectedOutput)
			c.Print(mockOutput)
		}

		DescribeTable("intCell", printTest,
			Entry("padded left", 1, 4, "   1"),
			Entry("not padded", 1234, 4, "1234"),
		)

		DescribeTable("stringCell", printTest,
			Entry("padded right", "hello", 10, "hello     "),
			Entry("no padding", "world", 5, "world"),
		)
	})
})
