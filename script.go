package jego

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

type Script interface {
	Run(ctx context.Context, arg any) (any, error)
	GetExport(name string) (any, error)
}

type script struct {
	vm *otto.Otto
	fn Value
	mr ModuleRegistery
}

func newScript(source io.Reader, mr ModuleRegistery) (Script, error) {
	s := &script{mr: mr}
	if err := s.init(source); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *script) init(source io.Reader) error {
	vm, err := createBaseVM(s.mr)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(source)
	if err != nil {
		return err
	}
	if _, err := vm.Run(string(b)); err != nil {
		return err
	}

	if s.fn, err = GetValue(vm, "exports", "default"); err != nil {
		return err
	}

	s.vm = vm
	return nil
}

func (s *script) Run(ctx context.Context, arg any) (any, error) {
	return s.fn.Call(otto.UndefinedValue(), arg)
}

func (s *script) GetExport(name string) (any, error) {
	v, err := GetValue(s.vm, "exports", name)
	if err != nil {
		return nil, err
	}
	return v.Export()
}
