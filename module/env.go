package module

import (
	"os"
	"strings"

	"github.com/ahmetcanozcan/jego/js"
	"github.com/robertkrimen/otto"
)

type env struct {
}

func Env() Module {
	return &env{}
}

func (e *env) Require(vm *otto.Otto) (any, error) {
	o := js.Object{}

	for _, e := range os.Environ() {
		parts := strings.Split(e, "=")
		o[parts[0]] = parts[1]
	}

	return o, nil
}

func (m *env) Copy() Module {
	return Env()
}
