package evolve

type command interface {
	Execute(backEnd BackEnd) error
}
