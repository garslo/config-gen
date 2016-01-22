package emitters

import (
	"fmt"
	"log"

	"github.com/garslo/config-gen/config"
)

type StdoutEmitter struct{}

func (me StdoutEmitter) Emit(state config.State) error {
	log.Printf("%d %d", len(state.Types), len(state.Decls))
	for _, t := range state.Types {
		fmt.Printf("%#v\n", t)
	}
	for _, d := range state.Decls {
		fmt.Printf("%#v\n", d)
	}
	return nil
}
