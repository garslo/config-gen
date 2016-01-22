package emitters

import (
	"encoding/json"
	"fmt"

	"github.com/garslo/config-gen/config"
)

type JsonEmitter struct{}

func (me JsonEmitter) Emit(state config.State) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
