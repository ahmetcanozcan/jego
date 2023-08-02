package jego

import (
	"fmt"

	"github.com/ahmetcanozcan/jego/js"
	"github.com/robertkrimen/otto"
)

func createBaseVM(mr ModuleRegistery) (VM, error) {
	vm := otto.New()
	vm.Set("require", func(name string) any {
		r, err := mr.Require(name)
		if err != nil {
			panic(err)
		}
		return r
	})
	err := runMultiScripts(
		vm,
		js.ExportJS,
	)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func runMultiScripts(vm VM, sc ...string) error {
	for _, s := range sc {
		if _, err := vm.Run(s); err != nil {
			return err
		}
	}
	return nil
}

func GetValue(vm VM, field string, nesteds ...string) (res Value, err error) {
	res, err = vm.Get(field)
	if err != nil {
		return res, err
	}

	for _, f := range nesteds {
		o := res.Object()
		if o == nil {
			return res, fmt.Errorf("%s is not an object", f)
		}
		res, err = o.Get(f)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func GetObject(vm VM, field string, nesteds ...string) (*Object, error) {
	v, err := GetValue(vm, field, nesteds...)
	if err != nil {
		return nil, err
	}
	return v.Object(), nil
}
