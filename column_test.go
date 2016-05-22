package gotable_test

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/tomcraven/gotable"
	. "github.com/tomcraven/gotable/gotable_mock"
)

var _ = Describe("Column", func() {
	It("creates a blank column", func() {
		NewColumn("", 0)
	})

	Describe("PrintHeader", func() {
		var (
			mockController *gomock.Controller
			mockOutput     *MockOutput
		)

		BeforeEach(func() {
			mockController = gomock.NewController(GinkgoT())
			mockOutput = NewMockOutput(mockController)
		})

		AfterEach(func() {
			mockController.Finish()
		})

		type headerFormat struct {
			name           string
			width          int
			expectedHeader string
		}

		DescribeTable("header formats",
			func(config headerFormat) {
				mockOutput.EXPECT().Print(config.expectedHeader)
				c := NewColumn(config.name, config.width)
				c.PrintHeader(mockOutput)
			},
			Entry("no padding", headerFormat{"test", 4, "test"}),
			Entry("pad right", headerFormat{"test", 5, "test "}),
			Entry("pad both", headerFormat{"test", 6, " test "}),
		)
	})

	Describe("push", func() {
		var (
			mockController *gomock.Controller
			mockOutput     *MockOutput
		)

		BeforeEach(func() {
			mockController = gomock.NewController(GinkgoT())
			mockOutput = NewMockOutput(mockController)
		})

		AfterEach(func() {
			mockController.Finish()
		})

		type rowPush struct {
			item           interface{}
			expectedOutput string
			printRowIndex  int
		}

		DescribeTable("different push types",
			func(config rowPush) {
				c := NewColumn("test", 10)
				mockOutput.EXPECT().Print(config.expectedOutput)
				c.Push(config.item)
				c.PrintCellAt(config.printRowIndex, mockOutput)
			},
			Entry("int", rowPush{
				item:           1,
				expectedOutput: "         1",
				printRowIndex:  0,
			}),
			Entry("string", rowPush{
				item:           "hello",
				expectedOutput: "hello     ",
				printRowIndex:  0,
			}),
		)
	})

	Describe("PrintCellAt", func() {
		It("panics if the ordinal is too high", func() {
			Expect(func() {
				c := NewColumn("test", 1)
				c.PrintCellAt(0, &NullOutput{}) // nothing inserted
			}).Should(Panic())
		})

		It("panics if the ordinal is negative", func() {
			Expect(func() {
				c := NewColumn("test", 1)
				c.PrintCellAt(-1, &NullOutput{})
			}).Should(Panic())
		})
	})
})
