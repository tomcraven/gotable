package gotable_test

import (
	"strconv"

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

		type headerFormatData struct {
			name           string
			width          int
			expectedHeader string
		}
		headerFormats := []headerFormatData{
			{"test", 4, "test"},
			{"test", 5, "test "},
			{"test", 6, " test "},
			{"a", 5, "  a  "},
		}

		for _, headerFormat := range headerFormats {
			Context("when the column's title is '"+headerFormat.name+"'", func() {
				Context("when the column's width is "+strconv.Itoa(headerFormat.width), func() {
					var c Column
					BeforeEach(func() {
						c = NewColumn(headerFormat.name, headerFormat.width)
					})

					It("prints the header correctly - '"+headerFormat.expectedHeader+"'", func() {
						mockOutput.EXPECT().Print(headerFormat.expectedHeader)
						c.PrintHeader(mockOutput)
					})
				})
			})
		}
	})
})
