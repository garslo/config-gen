package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type YamlFile string

func (me YamlFile) IsTypeFile() bool {
	return strings.HasSuffix(filepath.Dir(string(me)), "types")
}

func (me YamlFile) AsTypes() ([]Type, error) {
	t := []Type{}
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

func (me YamlFile) AsDecls() ([]Decl, error) {
	d := []Decl{}
	err := me.unmarshal(&d)
	return d, err
}

type YamlParser struct {
	Files []YamlFile
}

func (me YamlParser) Parse(root string) (ConfigState, error) {
	state := ConfigState{}
	if err := filepath.Walk(root, me.findYamlFiles); err != nil {
		return state, err
	}
	for _, file := range me.Files {
		if file.IsTypeFile() {
			ts, err := file.AsTypes()
			if err != nil {
				return state, err
			}
			if err := state.AddTypes(ts); err != nil {
				return state, err
			}
		} else {
			ds, err := file.AsDecls()
			if err != nil {
				return state, err
			}
			state.AddDecls(ds)
		}
	}
	return state, nil
}

func (me *YamlParser) findYamlFiles(path string, fileInfo os.FileInfo, err error) error {
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
