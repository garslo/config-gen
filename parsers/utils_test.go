package parsers_test

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/garslo/config-gen/config"
	"gopkg.in/yaml.v2"

	. "github.com/onsi/gomega"
)

func LayDownConfigTree(types []config.Type, decls []config.Decl) (root string) {
	var err error
	root, err = ioutil.TempDir("/tmp", "config-tree-")
	Expect(err).To(BeNil())
	typesDir := path.Join(root, "types")
	Expect(os.Mkdir(typesDir, os.ModeDir|0777)).To(Succeed())
	declsData, err := yaml.Marshal(decls)
	Expect(err).To(BeNil())
	declsFile := path.Join(root, "decls.yaml")
	Expect(ioutil.WriteFile(declsFile, declsData, 0666)).To(Succeed())

	typesData, err := yaml.Marshal(types)
	Expect(err).To(BeNil())
	typesFile := path.Join(typesDir, "types.yaml")
	Expect(ioutil.WriteFile(typesFile, typesData, 0666)).To(Succeed())
	return root
}
