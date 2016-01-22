package parsers_test

import (
	parsers "github.com/garslo/config-gen/parsers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yaml", func() {
	Context("the parser", func() {
		var (
			parser parsers.YamlParser
		)

		BeforeEach(func() {
			parser = parsers.YamlParser{}
		})

	})

	Context("the files", func() {
		Context("type files", func() {
			It("should recognize a type file", func() {
				file := parsers.YamlFile("/a/types/file.yaml")
				Expect(file.IsTypeFile()).To(BeTrue())
			})

			It("should recognize a file that's not a type file", func() {
				file := parsers.YamlFile("/just/a/normal/file.yaml")
				Expect(file.IsTypeFile()).To(BeFalse())
			})

			It("should extract a []config.Type from a type file", func() {

			})
		})
	})
})
