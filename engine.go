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
		mr: NewRegistery(),
	}
	if err := e.init(); err != nil {
		panic(err)
	}

	return e
}

func (e *Engine) init() error {
	return nil
}

func (e *Engine) SetRegistery(mr ModuleRegistery) *Engine {
	e.mr = mr
	return e
}

func (e *Engine) Script(buff io.Reader) (Script, error) {
	src, err := js.Transform(buff)
	if err != nil {
		return nil, err
	}
	return newScript(src, e.mr.Copy())
}
