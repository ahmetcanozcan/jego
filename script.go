package jego

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/ahmetcanozcan/jego/js"
	"github.com/ahmetcanozcan/jego/util"
	"github.com/robertkrimen/otto"
)

type Script interface {
	Run(ctx context.Context, arg any) (any, error)
	GetExport(name string) (any, error)
}

type script struct {
	vm *otto.Otto
	fn otto.Value
	mr *moduleRegistery
}

func newScript(source io.Reader, mr *moduleRegistery) (Script, error) {
	s := &script{mr: mr}
	if err := s.init(source); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *script) init(source io.Reader) error {
	s.vm = s.createBaseVM()

	b, err := ioutil.ReadAll(source)
	if err != nil {
		return err
	}
	if _, err := s.vm.Run(string(b)); err != nil {
		return err
	}

	if s.fn, err = util.GetValue(s.vm, "exports", "default"); err != nil {
		return err
	}

	return nil
}

func (s *script) Run(ctx context.Context, arg any) (any, error) {
	return s.fn.Call(otto.UndefinedValue(), arg)
}

func (s *script) GetExport(name string) (any, error) {
	v, err := util.GetValue(s.vm, "exports", name)
	if err != nil {
		return nil, err
	}
	return v.Export()
}

func (s *script) createBaseVM() *otto.Otto {
	vm := otto.New()
	vm.Set("require", s.require)
	util.RunMultiScripts(
		vm,
		js.ExportJS,
	)
	return vm
}

func (s *script) require(name string) any {
	r, err := s.mr.require(name)
	if err != nil {
		panic(err)
	}
	return r
}
