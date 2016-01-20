package main

import (
	"fmt"
	"log"
)

type StdoutEmitter struct{}

func (me StdoutEmitter) Emit(state ConfigState) error {
	log.Printf("%d %d", len(state.Types), len(state.Decls))
	for _, t := range state.Types {
		fmt.Printf("%#v\n", t)
	}
	for _, d := range state.Decls {
		fmt.Printf("%#v\n", d)
	}
	return nil
}
