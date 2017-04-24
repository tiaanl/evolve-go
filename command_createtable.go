package evolve

func NewCreateTableCommand(table Table) Command {
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
