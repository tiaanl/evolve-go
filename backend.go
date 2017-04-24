package evolve

type BackEndType int

type BackEnd interface {
	CreateTable(table Table) error
	DropTable(name string) error
}
