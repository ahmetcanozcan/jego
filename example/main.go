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
	mr := jego.NewRegistery(func(name string) (any, error) {
		return "default_value", nil
	})
	mr.Register("custom", &customModule{})
	mr.Register("echo", jego.ValueModule(func(v string) string {
		return "echo: " + v
	}))
	mr.Register("functional", jego.FuncModule(func() (any, error) {
		return jego.JSObject{
			"foo": "bar",
			"zoo": 18,
		}, nil
	}))

	modJS, err := jego.JSModule(strings.NewReader(modJS), mr)
	if err != nil {
		panic(err)
	}
	mr.Register("mod", modJS)

	e := jego.New().SetRegistery(mr)

	s, _ := e.Script(strings.NewReader(testFile))

	foo, _ := s.GetExport("foo")
	fmt.Println(foo) // prints bar

	r, _ := s.Run(context.Background(), 5)
	fmt.Println(r) // prints 10
	r, _ = s.Run(context.Background(), 15)
	fmt.Println(r) // prints 30
}
