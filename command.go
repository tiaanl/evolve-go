package evolve

type Command interface {
	Execute(backEnd BackEnd) error
}
