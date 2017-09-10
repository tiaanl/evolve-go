package evolve

func newDropTableCommand(tableName string) *dropTableCommand {
	return &dropTableCommand{
		tableName: tableName,
	}
}

type dropTableCommand struct {
	tableName string
}

func (c *dropTableCommand) ToSQL(dialect Dialect) (string, error) {
	return dialect.GetDropTableSQL(c.tableName)
}

func (c *dropTableCommand) Execute(backEnd BackEnd) error {
	return backEnd.DropTable(c.tableName)
}
