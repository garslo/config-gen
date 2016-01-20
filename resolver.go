package main

type Parser interface {
	Parse(root string) (ConfigState, error)
}

type Emitter interface {
	Emit(ConfigState) error
}

type Resolver struct {
	Parser  Parser
	Emitter Emitter
}

func (me Resolver) Emit(root string) error {
	state, err := me.Parser.Parse(root)
	if err != nil {
		return err
	}
	if err := state.Validate(); err != nil {
		return err
	}
	if err := me.Emitter.Emit(state); err != nil {
		return err
	}
	return nil
}
