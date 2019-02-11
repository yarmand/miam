package importer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// boo
var _ = Describe("ExtractInfo", func() {
	Describe("GetCreationDate", func() {
		It("extract the creation date properly", func() {
			fname := "../testAssets/testImage.jpeg"
			date := GetCreationDate(fname)
			Expect(date.Year()).To(Equal(2018))
			Expect(int(date.Month())).To(Equal(6))
			Expect(date.Day()).To(Equal(25))
		})
	})
})
