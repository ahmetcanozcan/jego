package jego

import (
	"io"

	"github.com/ahmetcanozcan/jego/js"
)

type Engine struct {
	mr ModuleRegistery
}

func New() *Engine {
	e := &Engine{
		mr: newModuleRegistery(),
	}
	if err := e.init(); err != nil {
		panic(err)
	}

	return e
}

func (e *Engine) init() error {
	return nil
}

func (e *Engine) Register(name string, module Module) *Engine {
	e.mr.Register(name, module)
	return e
}

func (e *Engine) Script(buff io.Reader) (Script, error) {
	src, err := js.Transform(buff)
	if err != nil {
		return nil, err
	}
	return newScript(src, e.mr.Copy())
}

func (e *Engine) ImportJSModule(src io.Reader) (Module, error) {
	transformed, err := js.Transform(src)
	if err != nil {
		return nil, err
	}
	sc, err := newScript(transformed, e.mr.Copy())
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
