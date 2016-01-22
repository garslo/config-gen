package parsers_test

import (
	"os"

	"github.com/garslo/config-gen/config"
	parsers "github.com/garslo/config-gen/parsers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yaml", func() {
	var (
		parser parsers.Yaml
	)

	BeforeEach(func() {
		parser = parsers.Yaml{}
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
		})

		Context("with actual files", func() {
			var (
				root  string
				types []config.Type
				decls []config.Decl
			)

			BeforeEach(func() {
				types = []config.Type{
					config.Type{Name: "t1"},
					config.Type{Name: "t2"},
					config.Type{Name: "t3"},
					config.Type{Name: "t4"},
				}
				decls = []config.Decl{
					config.Decl{TypeName: "t1", Name: "d1"},
					config.Decl{TypeName: "t2", Name: "d2"},
					config.Decl{TypeName: "t3", Name: "d3"},
					config.Decl{TypeName: "t4", Name: "d4"},
				}
				root = LayDownConfigTree(types, decls)
			})

			AfterEach(func() {
				Expect(os.RemoveAll(root)).To(Succeed())
			})

			It("should load the files", func() {
				_, err := parser.Parse(root)
				Expect(err).To(BeNil())
			})

			It("should load all the decls", func() {
				state, err := parser.Parse(root)
				Expect(err).To(BeNil())
				Expect(state.Decls).To(BeEquivalentTo(decls))
			})

			It("should load all the types", func() {
				state, err := parser.Parse(root)
				Expect(err).To(BeNil())
				names, parsedTypes := NamesAndTypes(state.Types)
				Expect(names).To(ConsistOf([]interface{}{"t1", "t2", "t3", "t4"}...))
				// An artifact of the yaml parser.
				parsedTypes = NillifyParamsIfEmpty(parsedTypes)
				Expect(parsedTypes).To(BeEquivalentTo(types))
			})
		})
	})
})

func NillifyParamsIfEmpty(types []config.Type) []config.Type {
	ret := []config.Type{}
	for _, t := range types {
		if len(t.Params) == 0 {
			t.Params = nil
		}
		ret = append(ret, t)
	}
	return ret
}

func NamesAndTypes(typesMap config.Types) ([]string, []config.Type) {
	types := []config.Type{}
	names := []string{}
	for name, t := range typesMap {
		types = append(types, t)
		names = append(names, name)
	}
	return names, types
}
