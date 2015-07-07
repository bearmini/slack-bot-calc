package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("Sample", func() {
	Context("top()", func() {
		It("return 200 Status", func() {
			RequestToRoot("GET", top)
			Expect(recorder.Code).To(Equal(200))
			Expect(recorder.Body).To(ContainSubstring("Hello"))
		})
	})
})

func TestSlackBotCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SlackBotCalc Suite")
}
