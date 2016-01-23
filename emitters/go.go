package emitters

import (
	"fmt"
	"go/format"
	"go/token"
	"os"
	"strings"

	"github.com/garslo/config-gen/config"
)

type Go struct{}

func (me Go) Emit(state config.State) error {
	file := File{
		Name: "config",
	}
	for _, t := range state.Types {
		fields := make([]Field, len(t.Params))
		for i, param := range t.Params {
			fields[i] = Field{
				Name: makeGoName(param.GoName, param.Name),
				Type: typeName(state.Types, param.Type),
				Tag:  fmt.Sprintf("`yaml:\"%s\"`", param.Name),
			}
		}
		s := Struct{
			Name:   makeGoName(t.GoName, t.Name),
			Fields: fields,
		}
		file.Structs = append(file.Structs, s)
	}

	fset := token.NewFileSet()
	return format.Node(os.Stdout, fset, file.Ast())
}

func typeName(types config.Types, name string) string {
	t, ok := types.Find(name)
	if ok {
		return makeGoName(t.GoName, t.Name)
	}
	return name
}

func makeGoName(def, name string) string {
	if def != "" {
		return def
	}
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return name
}
