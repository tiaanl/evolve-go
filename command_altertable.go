package evolve

func newAlterTableCommand(tableName string, atc *alterTableColumns) *alterTableCommand {
	return &alterTableCommand{
		tableName: tableName,
		atc:       atc,
	}
}

type alterTableCommand struct {
	tableName string
	atc       *alterTableColumns
}

func (c *alterTableCommand) ToSQL(dialect Dialect) (string, error) {
	return dialect.GetAlterTableSQL(c.tableName, c.atc)
}

func (c *alterTableCommand) Execute(backEnd BackEnd) error {
	return nil
}
