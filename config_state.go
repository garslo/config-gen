package main

import (
	"errors"
	"fmt"
)

type Param struct {
	Name     string `yaml:"name"`
	Short    string `yaml:"short"`
	Long     string `yaml:"long"`
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
}

type Type struct {
	Name   string  `yaml:"name"`
	Short  string  `yaml:"short"`
	Long   string  `yaml:"long"`
	GoName string  `yaml:"go_name"`
	Params []Param `yaml:"params"`
}

var (
	ErrTypeExists = errors.New("type already exists")
)

type Types map[string]Type

func (me Types) Find(name string) (Type, bool) {
	t, ok := me[name]
	return t, ok
}

type Decl struct {
	Name     string `yaml:"name"`
	TypeName string `yaml:"type"`
}

type Decls []Decl

type ConfigState struct {
	Types Types
	Decls Decls
}

func (me *ConfigState) AddTypes(ts []Type) error {
	for _, t := range ts {
		if me.Types == nil {
			me.Types = make(Types)
		}
		if _, ok := me.Types.Find(t.Name); ok {
			return ErrTypeExists
		}
		me.Types[t.Name] = t
	}
	return nil
}

func (me *ConfigState) AddDecls(ds []Decl) {
	me.Decls = append(me.Decls, ds...)
}

func (me ConfigState) Validate() error {
	for tName, t := range me.Types {
		for _, param := range t.Params {
			if _, ok := me.Types.Find(param.Type); !ok {
				return fmt.Errorf("Invalid type: could not find type '%s' for type '%s.%s'", param.Type, tName, param.Name)
			}
		}
	}
	for _, decl := range me.Decls {
		_, ok := me.Types.Find(decl.TypeName)
		if !ok {
			return fmt.Errorf("Invalid type: could not find type '%s' for decl '%s'", decl.TypeName, decl.Name)
		}
	}
	return nil
}
