package config_test

import (
	"fmt"

	. "github.com/garslo/config-gen/config"

	. "github.com/garslo/config-gen/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func MakeTypes(n int) []Type {
	ret := make([]Type, n)
	for i := 0; i < n; i++ {
		ret[i] = Type{
			Name: fmt.Sprintf("type%d", i),
		}
	}
	return ret
}

func MakeDecls(n int) []Decl {
	ret := make([]Decl, n)
	for i := 0; i < n; i++ {
		ret[i] = Decl{
			Name:     fmt.Sprintf("decl%d", i),
			TypeName: fmt.Sprintf("type%d", i*i),
		}
	}
	return ret
}

func MakeDeclsForTypes(types []Type) []Decl {
	decls := []Decl{}
	for i, t := range types {
		decls = append(decls, Decl{
			Name:     fmt.Sprintf("decl%d", i),
			TypeName: t.Name,
		})
	}
	return decls
}

var _ = Describe("State", func() {
	var (
		state State
		types []Type
		decls []Decl
	)

	BeforeEach(func() {
		state = State{}
	})

	Context("adding things", func() {
		BeforeEach(func() {
			types = MakeTypes(10)
			decls = MakeDecls(10)
		})

		It("should add a list of types", func() {
			state.AddTypes(types...)
			Expect(state.Types).To(HaveSameLen(types))
		})

		It("should store all the type names", func() {
			state.AddTypes(types...)
			for _, t := range types {
				Expect(state.Types).To(HaveKey(t.Name))
			}
		})

		It("should add a list of decls", func() {
			state.AddDecls(decls...)
			Expect(state.Decls).To(HaveSameLen(decls))
		})
	})

	Context("validating things", func() {
		Context("when everything is valid", func() {
			BeforeEach(func() {
				types = MakeTypes(5)
				decls = MakeDeclsForTypes(types)
				state.AddTypes(types...)
				state.AddDecls(decls...)
			})

			It("should pass validation", func() {
				Expect(state.Validate()).To(Succeed())
			})

			It("should pass validation with an empty state", func() {
				state = State{}
				Expect(state.Validate()).To(Succeed())
			})
		})

		Context("invalid setup", func() {
			It("should fail when a decl has a non-existing type", func() {
				state.AddDecls(Decl{
					TypeName: "i do not exist",
				})
				Expect(state.Validate()).NotTo(Succeed())
			})

			It("should fail when a type does not have a name", func() {
				state.AddTypes(Type{})
				Expect(state.Validate()).NotTo(Succeed())
			})

			It("should fail when a decl has no name", func() {
				t := Type{
					Name: "testtype",
				}
				state.AddTypes(t)
				state.AddDecls(Decl{
					TypeName: t.Name,
				})
				Expect(state.Validate()).NotTo(Succeed())
			})
		})
	})
})
