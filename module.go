package jego

import (
	"io"

	"github.com/ahmetcanozcan/jego/js"
)

type Module interface {
	Require() (any, error)
	Copy() Module
}

var _ Module = (FuncModule)(nil)

type FuncModule func() (any, error)

func (fm FuncModule) Require() (any, error) {
	return fm()
}

func (fm FuncModule) Copy() Module {
	return fm
}

type valueModule struct {
	val any
}

func ValueModule(v any) Module {
	return &valueModule{v}
}

func (m *valueModule) Require() (any, error) {
	return m.val, nil
}

func (m *valueModule) Copy() Module {
	return m
}

func JSModule(src io.Reader, registery ...ModuleRegistery) (Module, error) {
	if len(registery) == 0 {
		registery = append(registery, NewRegistery())
	}
	mr := registery[0]

	transformed, err := js.Transform(src)
	if err != nil {
		return nil, err
	}
	sc, err := newScript(transformed, mr.Copy())
	if err != nil {
		return nil, err
	}

	v, err := sc.GetValue("default")
	if err != nil {
		return nil, err
	}

	if v.IsObject() {
		return ValueModule(v.Object()), nil

	}

	return ValueModule(v), nil
}
