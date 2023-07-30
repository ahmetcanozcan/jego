package module

import (
	"os"
	"strings"

	"github.com/ahmetcanozcan/jego/js"
)

type env struct {
}

func Env() Module {
	return &env{}
}

func (e *env) Require() (any, error) {
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
