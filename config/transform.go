package config

type Parser interface {
	Parse(string) (State, error)
}

type Emitter interface {
	Emit(State) error
}

func Transform(root string, parser Parser, emitter Emitter) error {
	state, err := parser.Parse(root)
	if err != nil {
		return err
	}
	if err := state.Validate(); err != nil {
		return err
	}
	if err := emitter.Emit(state); err != nil {
		return err
	}
	return nil
}
