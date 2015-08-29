package numeroil_test

import (
	. "github.com/dgodd/numeroil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Numeroil", func() {
	Describe("AddLetters", func() {
		It("lavender returns 36", func() {
			Expect(AddLetters("lavender")).To(Equal(36))
		})

		It("LAVENDER returns 36", func() {
			Expect(AddLetters("lavender")).To(Equal(36))
		})
	})

	Describe("Reduce", func() {
		It("3 returns 3", func() {
			Expect(Reduce(3)).To(Equal(3))
		})

		It("36 returns 9", func() {
			Expect(Reduce(36)).To(Equal(9))
		})

		It("98 returns 8", func() {
			Expect(Reduce(98)).To(Equal(8))
		})

		It("11 is special, thus 56 returns 11", func() {
			Expect(Reduce(56)).To(Equal(11))
		})

		It("22 is special, thus 22 returns 22", func() {
			Expect(Reduce(22)).To(Equal(22))
		})

		It("33 is special, thus 33 returns 33", func() {
			Expect(Reduce(33)).To(Equal(33))
		})

		It("44 is special, thus 44 returns 44", func() {
			Expect(Reduce(44)).To(Equal(44))
		})

		It("55 is NOT  special, thus 55 returns 1", func() {
			Expect(Reduce(55)).To(Equal(1))
		})
	})
})
