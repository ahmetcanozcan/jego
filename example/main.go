package main

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/ahmetcanozcan/jego"
)

//go:embed file.js
var testFile string

//go:embed mod.js
var modJS string

type customModule struct {
}

func (*customModule) Require() (any, error) {
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

	e := jego.New()
	e.Register("custom", &customModule{})
	e.Register("echo", jego.ValueModule(func(v string) string {
		return "echo: " + v
	}))
	e.Register("functional", jego.FuncModule(func() (any, error) {
		return jego.JSObject{
			"foo": "bar",
			"zoo": 18,
		}, nil
	}))

	modJS, err := e.ImportJSModule(strings.NewReader(modJS))
	if err != nil {
		panic(err)
	}

	e.Register("mod", modJS)

	s, _ := e.Script(strings.NewReader(testFile))

	foo, _ := s.GetExport("foo")
	fmt.Println(foo) // prints bar

	r, _ := s.Run(context.Background(), 5)
	fmt.Println(r) // prints 10
	r, _ = s.Run(context.Background(), 15)
	fmt.Println(r) // prints 30
}
