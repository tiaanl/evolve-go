package evolve

func NewDropTableCommand(tableName string) Command {
	return &dropTableCommand{
		tableName: tableName,
	}
}

type dropTableCommand struct {
	tableName string
}

func (c *dropTableCommand) Execute(backEnd BackEnd) error {
	return backEnd.DropTable(c.tableName)
}
