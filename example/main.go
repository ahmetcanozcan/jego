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

func main() {
	e := jego.New()
	s, _ := e.Script(strings.NewReader(testFile))
	r, _ := s.Run(context.Background(), 5)
	fmt.Println(r) // prints 10
}
