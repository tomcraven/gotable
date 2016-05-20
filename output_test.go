package gotable_test

import (
	"strings"

	gomock "github.com/golang/mock/gomock"
	. "github.com/tomcraven/gotable"
	. "github.com/tomcraven/gotable/gotable_mock"
)

var _ = Describe("Output", func() {
	Describe("NullOutput", func() {
		It("does nothing when printing", func() {
			output := NullOutput{}
			output.Print("noop")
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
