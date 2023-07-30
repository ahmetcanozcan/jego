package module

type Module interface {
	Require() (any, error)
	Copy() Module
}
