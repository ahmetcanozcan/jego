package module

import "github.com/robertkrimen/otto"

type Module interface {
	Require(vm *otto.Otto) (any, error)
	Copy() Module
}
