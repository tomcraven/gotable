package gotable_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/tomcraven/gotable"
	. "github.com/tomcraven/gotable/gotable_mock"
)

var _ = Describe("Table", func() {
	It("creates a blank table", func() {
		_ = NewTable([]Column{})
	})

	Describe("Push", func() {
		It("pushes a row into the table", func() {

		})
	})

	Describe("Print", func() {
		type tableConfiguration struct {
			columns        []Column
			expectedOutput []string
		}

		var tableConfigurations []tableConfiguration

		printTest := func() {
			var (
				t              Table
				mockController *gomock.Controller
				mockOutput     *MockOutput
			)

			for _, tableConfig := range tableConfigurations {
				BeforeEach(func() {
					t = NewTable(tableConfig.columns)
					mockController = gomock.NewController(GinkgoT())
					mockOutput = NewMockOutput(mockController)
				})

				It("prints correctly", func() {
					for _, expectedOutput := range tableConfig.expectedOutput {
						mockOutput.EXPECT().Print(expectedOutput)
					}
					t.Print(mockOutput)
				})

				AfterEach(func() {
					mockController.Finish()
				})
			}
		}

		Context("when a table has a single column", func() {
			tableConfigurations = []tableConfiguration{
				{
					columns: []Column{
						NewColumn("test", 4),
					},
					expectedOutput: []string{
						"+----+",
						"|test|",
						"+----+",
					},
				},
				{
					columns: []Column{
						NewColumn("another test", 20),
					},
					expectedOutput: []string{
						"+--------------------+",
						"|    another test    |",
						"+--------------------+",
					},
				},
				{
					columns: []Column{
						NewColumn("test odd", 9),
					},
					expectedOutput: []string{
						"+---------+",
						"|test odd |",
						"+---------+",
					},
				},
			}

			printTest()
		})

		Context("when a table has multiple columns", func() {
			tableConfigurations = []tableConfiguration{
				{
					columns: []Column{
						NewColumn("test", 4),
						NewColumn("test", 4),
					},
					expectedOutput: []string{
						"+----+----+",
						"|test|test|",
						"+----+----+",
					},
				},
				{
					columns: []Column{
						NewColumn("col", 5),
						NewColumn("middle column", 20),
						NewColumn("c", 1),
					},
					expectedOutput: []string{
						"+-----+--------------------+-+",
						"| col |   middle column    |c|",
						"+-----+--------------------+-+",
					},
				},
			}

			printTest()
		})
	})
})
