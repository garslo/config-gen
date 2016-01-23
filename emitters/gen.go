package emitters

import (
	"go/ast"
	"go/token"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

func (me Field) Ast() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			&ast.Ident{
				Name: me.Name,
				Obj: &ast.Object{
					Kind: ast.Var,
					Name: me.Name,
				},
			},
		},
		Type: &ast.Ident{
			Name: me.Type,
		},
		Tag: &ast.BasicLit{
			Kind:  token.STRING,
			Value: me.Tag,
		},
	}
}

type Struct struct {
	Name   string
	Fields []Field
}

func (me Struct) Ast() *ast.TypeSpec {
	fields := make([]*ast.Field, len(me.Fields))
	for i, field := range me.Fields {
		fields[i] = field.Ast()
	}
	return &ast.TypeSpec{
		Name: &ast.Ident{
			Name: me.Name,
			Obj: &ast.Object{
				Kind: ast.Typ,
				Name: me.Name,
			},
		},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: fields,
			},
		},
	}
}

type File struct {
	Name    string
	Structs []Struct
}

func (me File) Ast() *ast.File {
	decls := make([]ast.Decl, len(me.Structs))
	for i, s := range me.Structs {
		decls[i] = &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{s.Ast()},
		}
	}
	return &ast.File{
		Name: &ast.Ident{
			Name: me.Name,
		},
		Decls: decls,
	}
}
