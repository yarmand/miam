package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configuration", func() {
	var data = `
importers:
  - source: "default"
    destination: "d:/lapin"
  - source: "e:/"
    destination: "d:/lapin"
`
	BeforeSuite(func() {
		fmt.Printf("== DATA ==\n%v\n===========", data)
	})
	It("load the config", func() {
		conf := parseConfig(data)
		for _, importer := range conf.Importers {
			Expect(importer.Destination).To(Equal("d:/lapin"))
		}
	})
})
