package module

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
