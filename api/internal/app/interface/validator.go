package app

type Validator interface {
	Struct(any) error
	Var(any, string) error
}
