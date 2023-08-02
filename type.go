package jego

import (
	"github.com/ahmetcanozcan/jego/js"
	"github.com/robertkrimen/otto"
)

type (
	JSObject = js.Object
	VM       = *otto.Otto
	Object   = otto.Object
	Value    = otto.Value
)
