package babble_test

import (
	. "github.com/amoschu/babble"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("babble", func() {
	var babbler Babbler
	BeforeEach(func() {
		babbler = Babbler{
			Words:     []string{"hello's"},
			Separator: "☃",
		}
	})

	It("returns a random word", func() {
		Expect(babbler.Babble(1)).To(Equal("hellos"))
	})

	It("concatenates strings", func() {
		Expect(babbler.Babble(2)).To(Equal("hellos☃hellos"))
	})

	It("sanitizes punctuation", func() {
		Expect(babbler.Babble(4)).To(Equal("hellos☃hellos☃hellos☃hellos"))
	})
})
