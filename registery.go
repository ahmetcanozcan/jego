package jego

import (
	"fmt"
	"sync"

	"github.com/ahmetcanozcan/jego/module"
	"github.com/robertkrimen/otto"
)

type ModuleRegistery interface {
	Register(name string, mod module.Module)
	Copy() ModuleRegistery
	Require(name string) (any, error)
}

type moduleRegistery struct {
	rw      sync.Mutex
	modules map[string]module.Module
	cacheRw sync.Mutex
	cache   map[string]*otto.Object
}

func newModuleRegistery() ModuleRegistery {
	return &moduleRegistery{
		modules: make(map[string]module.Module),
		cache:   make(map[string]*otto.Object),
	}
}

func (mr *moduleRegistery) Register(name string, mod module.Module) {
	mr.rw.Lock()
	defer mr.rw.Unlock()
	mr.modules[name] = mod
}

func (mr *moduleRegistery) copyModules() map[string]module.Module {
	mr.rw.Lock()
	defer mr.rw.Unlock()
	cp := make(map[string]module.Module, len(mr.modules))
	for k, v := range mr.modules {
		cp[k] = v.Copy()
	}
	return cp
}

func (mr *moduleRegistery) Copy() ModuleRegistery {
	return &moduleRegistery{
		modules: mr.copyModules(),
		cache:   make(map[string]*otto.Object),
	}
}

func (mr *moduleRegistery) Require(name string) (any, error) {
	mr.rw.Lock()
	defer mr.rw.Unlock()
	m, ok := mr.modules[name]
	if !ok {
		return nil, fmt.Errorf("module %s not found!", name)
	}

	return m.Require()
}
