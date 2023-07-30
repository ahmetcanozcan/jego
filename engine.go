package jego

import (
	"io"

	"github.com/ahmetcanozcan/jego/js"
	"github.com/ahmetcanozcan/jego/module"
)

type Engine struct {
	mr *moduleRegistery
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
	e.mr.register("env", module.Env())
	return nil
}

func (e *Engine) Register(name string, module module.Module) *Engine {
	e.mr.register(name, module)
	return e
}

func (e *Engine) Script(buff io.Reader) (Script, error) {
	src, err := js.Transform(buff)
	if err != nil {
		return nil, err
	}
	return newScript(src, e.mr.copy())
}
