package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configuration", func() {
	var data = `
importers:
	- source: default
		destination: "d:/lapin"
`
	BeforeEach(func() {
		fmt.Println(data)
	})
	It("load the config", func() {
		parseConfig(data)
		Expect(1).To(Equal(1))
	})
})
