package evolve

func newCreateTableCommand(table Table) command {
	return &createTableCommand{
		table: table,
	}
}

type createTableCommand struct {
	table Table
}

func (c *createTableCommand) ToSQL(dialect Dialect) (string, error) {
	return dialect.GetCreateTableSQL(c.table)
}

func (c *createTableCommand) Execute(backEnd BackEnd) error {
	return backEnd.CreateTable(c.table)
}
