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

		printTest := func(configs []printConfiguration) {
			for _, config := range configs {
				It("prints the cell correctly", func() {
					c = NewCell(mockColumn, config.input)
					mockColumn.EXPECT().GetWidth().Return(config.width)
					mockOutput.EXPECT().Print(config.expectedOutput)
					c.Print(mockOutput)
				})
			}
		}

		Describe("intCell", func() {
			printTest([]printConfiguration{
				{1, 4, "   1"},
				{1234, 4, "1234"},
			})
		})

		Describe("stringCell", func() {
			printTest([]printConfiguration{
				{"hello", 10, "hello     "},
				{"world", 5, "world"},
			})
		})
	})
})
