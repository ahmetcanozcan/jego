package jego

type Module interface {
	Require() (any, error)
	Copy() Module
}

var _ Module = (FuncModule)(nil)

type FuncModule func() (any, error)

func (fm FuncModule) Require() (any, error) {
	return fm()
}

func (fm FuncModule) Copy() Module {
	return fm
}

type valueModule struct {
	val any
}

func ValueModule(v any) Module {
	return &valueModule{v}
}

func (m *valueModule) Require() (any, error) {
	return m.val, nil
}

func (m *valueModule) Copy() Module {
	return m
}
