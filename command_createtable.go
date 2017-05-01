package evolve

func newCreateTableCommand(table Table) command {
	return &createTableCommand{
		table: table,
	}
}

type createTableCommand struct {
	table Table
}

func (c *createTableCommand) Execute(backEnd BackEnd) error {
	return backEnd.CreateTable(c.table)
}
