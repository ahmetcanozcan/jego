package util

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func RunMultiScripts(vm *otto.Otto, sc ...string) error {
	for _, s := range sc {
		if _, err := vm.Run(s); err != nil {
			return err
		}
	}
	return nil
}

func GetValue(vm *otto.Otto, field string, nesteds ...string) (res otto.Value, err error) {
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

func GetObject(vm *otto.Otto, field string, nesteds ...string) (*otto.Object, error) {
	v, err := GetValue(vm, field, nesteds...)
	if err != nil {
		return nil, err
	}
	return v.Object(), nil
}
