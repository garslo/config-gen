package emitters

import (
	"fmt"
	"os"
	"strings"

	"github.com/garslo/config-gen/config"

	g "github.com/garslo/gogen"
)

type Go struct{}

func (me Go) Emit(state config.State) error {
	pkg := g.Package{
		Name: "config",
	}
	for _, t := range state.Types {
		fields := make([]g.Field, len(t.Params))
		for i, param := range t.Params {
			fields[i] = g.Field{
				Name:     makeGoName(param.GoName, param.Name),
				TypeName: typeName(state.Types, param.Type),
				Tag:      fmt.Sprintf("`yaml:\"%s\"`", param.Name),
			}
		}
		pkg.Declare(g.Struct{
			Name:   makeGoName(t.GoName, t.Name),
			Fields: fields,
		})
	}
	return pkg.WriteTo(os.Stdout)
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
