package main

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/ahmetcanozcan/jego"
	"github.com/robertkrimen/otto"
)

//go:embed file.js
var testFile string

type customModule struct {
}

func (*customModule) Require(*otto.Otto) (any, error) {
	return jego.JSObject{
		"double": func(b float64) float64 {
			return b * 2
		},
	}, nil
}

func (cm *customModule) Copy() jego.Module {
	return cm
}

func main() {
	e := jego.New().Register("custom", &customModule{})
	s, _ := e.Script(strings.NewReader(testFile))
	r, _ := s.Run(context.Background(), 5)
	fmt.Println(r) // prints 10
	r, _ = s.Run(context.Background(), 15)
	fmt.Println(r) // prints 30
}
