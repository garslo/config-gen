package parsers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/garslo/config-gen/config"
	"gopkg.in/yaml.v2"
)

type YamlFile string

func (me YamlFile) IsTypeFile() bool {
	return strings.HasSuffix(filepath.Dir(string(me)), "types")
}

func (me YamlFile) AsTypes() ([]config.Type, error) {
	t := []config.Type{}
	err := me.unmarshal(&t)
	return t, err
}

func (me YamlFile) unmarshal(out interface{}) error {
	data, err := ioutil.ReadFile(string(me))
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}

func (me YamlFile) AsDecls() ([]config.Decl, error) {
	d := []config.Decl{}
	err := me.unmarshal(&d)
	return d, err
}

type Yaml struct {
	Files []YamlFile
}

func (me Yaml) Parse(root string) (config.State, error) {
	state := config.State{}
	if err := filepath.Walk(root, me.findYamlFiles); err != nil {
		return state, err
	}
	for _, file := range me.Files {
		if file.IsTypeFile() {
			ts, err := file.AsTypes()
			if err != nil {
				return state, err
			}
			if err := state.AddTypes(ts...); err != nil {
				return state, err
			}
		} else {
			ds, err := file.AsDecls()
			if err != nil {
				return state, err
			}
			state.AddDecls(ds...)
		}
	}
	return state, nil
}

func (me *Yaml) findYamlFiles(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		return nil
	}
	ext := filepath.Ext(path)
	if ext == ".yaml" || ext == ".yml" {
		me.Files = append(me.Files, YamlFile(path))
	}
	return nil
}
