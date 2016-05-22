package gotable_test

import (
	"strings"

	gomock "github.com/golang/mock/gomock"
	. "github.com/tomcraven/gotable"
	. "github.com/tomcraven/gotable/gotable_mock"
)

var _ = Describe("Output", func() {
	Describe("NullOutput", func() {
		var output NullOutput

		BeforeEach(func() {
			output = NullOutput{}
		})

		It("does nothing when printing", func() {
			output.Print("noop")
		})

		It("does nothing when flushing", func() {
			output.Flush()
		})
	})

	Describe("OutputStdOut", func() {
		var output OutputStdOut

		BeforeEach(func() {
			output = OutputStdOut{}
		})

		It("prints to stdout", func() {
			output.Print("test")
		})

		It("does nothing when flushing", func() {
			output.Flush()
		})
	})

	Describe("OutputAligned", func() {
		var (
			mockOutput     *MockOutput
			mockController *gomock.Controller
		)

		BeforeEach(func() {
			mockController = gomock.NewController(GinkgoT())
			mockOutput = NewMockOutput(mockController)
		})

		AfterEach(func() {
			mockController.Finish()
		})

		type config struct {
			fixedWidth     int
			inputCallback  func(*OutputAligned)
			expectedOutput string
			alignment      Alignment
		}

		DescribeTable("print",
			func(config config) {
				output := NewOutputAligned(config.fixedWidth, mockOutput, config.alignment)
				mockOutput.EXPECT().Print(config.expectedOutput)
				config.inputCallback(&output)
			},
			Entry("blank, 0 width", config{
				fixedWidth:     0,
				inputCallback:  func(o *OutputAligned) { o.Print("") },
				expectedOutput: "",
			}),
			Entry("blank, 5 width", config{
				fixedWidth:     5,
				inputCallback:  func(o *OutputAligned) { o.Print("") },
				expectedOutput: "     ",
			}),
			Entry("'test' with no pad needed", config{
				fixedWidth:     4,
				inputCallback:  func(o *OutputAligned) { o.Print("test") },
				expectedOutput: "test",
			}),
			Entry("'test' with padding needed align left", config{
				fixedWidth:     10,
				inputCallback:  func(o *OutputAligned) { o.Print("test") },
				expectedOutput: "test      ",
				alignment:      Left,
			}),
			Entry("'test' with padding needed align right", config{
				fixedWidth:     10,
				inputCallback:  func(o *OutputAligned) { o.Print("test") },
				expectedOutput: "      test",
				alignment:      Right,
			}),
			Entry("'test' with truncation needed", config{
				fixedWidth:     2,
				inputCallback:  func(o *OutputAligned) { o.Print("test") },
				expectedOutput: "te",
			}),
		)

		It("flushes", func() {
			output := NewOutputAligned(1, &NullOutput{}, Left)
			output.Flush()
		})
	})

	Describe("OutputBuffered", func() {
		type bufferedOutputData struct {
			inputs         []string
			expectedOutput string
		}
		bufferedOutputTestData := []bufferedOutputData{
			{
				[]string{"abc"},
				"abc",
			},
			{
				[]string{"d", "e", "f"},
				"def",
			},
		}

		Context("when creating a new OutputBuffered for each test", func() {
			for _, data := range bufferedOutputTestData {
				Context("when the input is '"+strings.Join(data.inputs, ", ")+"'", func() {
					var (
						mockController *gomock.Controller
						mockOutput     *MockOutput
						outputBuffered OutputBuffered
					)

					BeforeEach(func() {
						mockController = gomock.NewController(GinkgoT())
						mockOutput = NewMockOutput(mockController)
						outputBuffered = NewOutputBuffered(mockOutput)

						for _, input := range data.inputs {
							outputBuffered.Print(input)
						}
					})

					It("sends '"+data.expectedOutput+"'to the consumer", func() {
						mockOutput.EXPECT().Print(data.expectedOutput)
						outputBuffered.Flush()
					})

					AfterEach(func() {
						mockController.Finish()
					})
				})
			}
		})

		Context("when reusing the same OutputBuffered for each test", func() {
			mockController := gomock.NewController(GinkgoT())
			mockOutput := NewMockOutput(mockController)
			outputBuffered := NewOutputBuffered(mockOutput)

			for _, data := range bufferedOutputTestData {
				Context("when the input is '"+strings.Join(data.inputs, ", ")+"'", func() {
					BeforeEach(func() {
						for _, input := range data.inputs {
							outputBuffered.Print(input)
						}
					})

					It("sends '"+data.expectedOutput+"'to the consumer", func() {
						mockOutput.EXPECT().Print(data.expectedOutput)
						outputBuffered.Flush()
					})
				})
			}

			mockController.Finish()
		})
	})
})
