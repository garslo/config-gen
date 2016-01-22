package config

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

func (me Type) Validate() error {
	if me.Name == "" {
		return errors.New("type has no name")
	}
	return nil
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

func (me Decl) Validate() error {
	if me.Name == "" {
		return errors.New("decl has no name")
	}
	return nil
}

type Decls []Decl

type State struct {
	Types Types
	Decls Decls
}

func (me *State) AddTypes(ts ...Type) error {
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

func (me *State) AddDecls(ds ...Decl) {
	me.Decls = append(me.Decls, ds...)
}

func (me State) Validate() error {
	for tName, t := range me.Types {
		if err := t.Validate(); err != nil {
			return fmt.Errorf("Invalid type '%s': %s", tName, err)
		}
		for _, param := range t.Params {
			if _, ok := me.Types.Find(param.Type); !ok {
				return fmt.Errorf("Invalid type: could not find type '%s' for type '%s.%s'", param.Type, tName, param.Name)
			}
		}
	}
	for _, decl := range me.Decls {
		if err := decl.Validate(); err != nil {
			return fmt.Errorf("Invalid decl '%s': %s", decl.Name, err)
		}
		_, ok := me.Types.Find(decl.TypeName)
		if !ok {
			return fmt.Errorf("Invalid type: could not find type '%s' for decl '%s'", decl.TypeName, decl.Name)
		}
	}
	return nil
}
