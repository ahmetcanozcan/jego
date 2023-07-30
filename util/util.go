package util

import "github.com/robertkrimen/otto"

func RunMultiScripts(vm *otto.Otto, sc ...string) error {
	for _, s := range sc {
		if _, err := vm.Run(s); err != nil {
			return err
		}
	}
	return nil
}

func GetObject(vm *otto.Otto, field string, nesteds ...string) (*otto.Object, error) {
	v, err := vm.Get(field)
	if err != nil {
		return nil, err
	}

	o := v.Object()
	for _, f := range nesteds {
		v, err := o.Get(f)
		if err != nil {
			return nil, err
		}
		o = v.Object()
	}

	return o, nil
}
