package jego

import (
	"fmt"
	"sync"
)

type ModuleRegistery interface {
	Register(name string, mod Module)
	Copy() ModuleRegistery
	Require(name string) (any, error)
}

type moduleRegistery struct {
	rw      sync.Mutex
	modules map[string]Module
	cacheRw sync.Mutex
	cache   map[string]*Object
}

func NewRegistery() ModuleRegistery {
	return &moduleRegistery{
		modules: make(map[string]Module),
		cache:   make(map[string]*Object),
	}
}

func (mr *moduleRegistery) Register(name string, mod Module) {
	mr.rw.Lock()
	defer mr.rw.Unlock()
	mr.modules[name] = mod
}

func (mr *moduleRegistery) copyModules() map[string]Module {
	mr.rw.Lock()
	defer mr.rw.Unlock()
	cp := make(map[string]Module, len(mr.modules))
	for k, v := range mr.modules {
		cp[k] = v.Copy()
	}
	return cp
}

func (mr *moduleRegistery) Copy() ModuleRegistery {
	return &moduleRegistery{
		modules: mr.copyModules(),
		cache:   make(map[string]*Object),
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
