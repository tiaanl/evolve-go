package evolve

type command interface {
	ToSQL(dialect Dialect) (string, error)
	Execute(backEnd BackEnd) error
}
