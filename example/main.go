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
	e := jego.New().
		Register("custom", &customModule{}).
		Register("echo", jego.ValueModule(func(v string) string {
			return "echo: " + v
		})).
		Register("functional", jego.FuncModule(func() (any, error) {
			return jego.JSObject{
				"foo": "bar",
				"zoo": 18,
			}, nil
		}))

	s, _ := e.Script(strings.NewReader(testFile))

	foo, _ := s.GetExport("foo")
	fmt.Println(foo) // prints bar

	r, _ := s.Run(context.Background(), 5)
	fmt.Println(r) // prints 10
	r, _ = s.Run(context.Background(), 15)
	fmt.Println(r) // prints 30
}
