package main

import (
	"encoding/json"
	"fmt"
)

type JsonEmitter struct{}

func (me JsonEmitter) Emit(state ConfigState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
