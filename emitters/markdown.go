package emitters

import (
	"fmt"
	"io"
	"os"

	"github.com/garslo/config-gen/config"
	"github.com/garslo/genmarkdown"
)

func wrap(in string) string {
	if in == "" {
		return in
	}
	return fmt.Sprintf("`%s`", in)
}

type Markdown struct{}

func (me Markdown) Emit(state config.State) error {
	nodes := genmarkdown.Nodes{}
	for _, t := range state.Types {
		nodes.Add(genmarkdown.Heading{
			Level: 1,
			Text:  t.Name,
		})
		if t.Long != "" {
			nodes.Add(genmarkdown.Paragraph{
				Text: t.Long,
			})
		}
		rows := make([]genmarkdown.Row, len(t.Params))
		for i, param := range t.Params {
			rows[i] = genmarkdown.Row{
				[]string{wrap(param.Name), wrap(param.Type), wrap(param.Default), param.Short},
			}
		}
		nodes.Add(genmarkdown.Table{
			Columns: []string{"Name", "Type", "Default", "Description"},
			Rows:    rows,
		})
	}
	_, err := io.Copy(os.Stdout, nodes.Markdown())
	return err
}
