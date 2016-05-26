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
		type tableConfiguration struct {
			callback       func(*Table)
			expectedOutput []string
		}

		var (
			t              Table
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

		printTest := func(config tableConfiguration) {
			config.callback(&t)
			for _, expectedOutput := range config.expectedOutput {
				mockOutput.EXPECT().Print(expectedOutput)
			}
			t.PrintTo(mockOutput)
		}

		Context("when the table has a single column", func() {
			BeforeEach(func() {
				t = NewTable([]Column{
					NewColumn("Test", 4),
				})
			})

			DescribeTable("pushes a row into the table", printTest,
				Entry("single row", tableConfiguration{
					callback: func(t *Table) {
						t.Push(1)
					},
					expectedOutput: []string{
						"┌──────┐",
						"│ Test │",
						"├──────┤",
						"│    1 │",
						"└──────┘",
					},
				}),
				Entry("multiple rows", tableConfiguration{
					callback: func(t *Table) {
						t.Push(42)
						t.Push(2323)
					},
					expectedOutput: []string{
						"┌──────┐",
						"│ Test │",
						"├──────┤",
						"│   42 │",
						"│ 2323 │",
						"└──────┘",
					},
				}),
			)
		})

		Context("when the table has multiple columns", func() {

			BeforeEach(func() {
				t = NewTable([]Column{
					NewColumn("col1", 6),
					NewColumn("col2", 10),
				})
			})

			DescribeTable("pushes a row into the table", printTest,
				Entry("single row", tableConfiguration{
					callback: func(t *Table) {
						t.Push(1, 2)
					},
					expectedOutput: []string{
						"┌────────┬────────────┐",
						"│  col1  │    col2    │",
						"├────────┼────────────┤",
						"│      1 │          2 │",
						"└────────┴────────────┘",
					},
				}),
				Entry("multiple rows", tableConfiguration{
					callback: func(t *Table) {
						t.Push(42, 23)
						t.Push(1234, 1234567)
					},
					expectedOutput: []string{
						"┌────────┬────────────┐",
						"│  col1  │    col2    │",
						"├────────┼────────────┤",
						"│     42 │         23 │",
						"│   1234 │    1234567 │",
						"└────────┴────────────┘",
					},
				}),
			)
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
					t.PrintTo(mockOutput)
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
						"┌──────┐",
						"│ test │",
						"├──────┤",
						"└──────┘",
					},
				},
				{
					columns: []Column{
						NewColumn("another test", 20),
					},
					expectedOutput: []string{
						"┌──────────────────────┐",
						"│     another test     │",
						"├──────────────────────┤",
						"└──────────────────────┘",
					},
				},
				{
					columns: []Column{
						NewColumn("test odd", 9),
					},
					expectedOutput: []string{
						"┌───────────┐",
						"│ test odd  │",
						"├───────────┤",
						"└───────────┘",
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
						"┌──────┬──────┐",
						"│ test │ test │",
						"├──────┼──────┤",
						"└──────┴──────┘",
					},
				},
				{
					columns: []Column{
						NewColumn("col", 5),
						NewColumn("middle column", 20),
						NewColumn("c", 1),
					},
					expectedOutput: []string{
						"┌───────┬──────────────────────┬───┐",
						"│  col  │    middle column     │ c │",
						"├───────┼──────────────────────┼───┤",
						"└───────┴──────────────────────┴───┘",
					},
				},
			}

			printTest()
		})
	})
})
