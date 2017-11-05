package evolve

import "strings"

type ChangeSet interface {
	ToSQL(dialect Dialect) (string, error)

	CreateTable(table Table)
	CreateTableWithColumns(tableName string, columns ...*Column)
	CreateTableWithFunc(tableName string, fn CreateTableFunc)
	DropTable(tableName string)
	AlterTable(tableName string, atc *alterTableColumns)

	Execute(backEnd BackEnd) error
}

func NewChangeSet() ChangeSet {
	return &changeSet{
		commandBus: newCommandBus(),
	}
}

func NewChangeSetWithCommandBus(commandBus *commandBus) ChangeSet {
	return &changeSet{
		commandBus: commandBus,
	}
}

type changeSet struct {
	commandBus *commandBus
}

func (cs *changeSet) ToSQL(dialect Dialect) (string, error) {
	lines := []string{}

	for _, command := range cs.commandBus.commands {
		query, err := command.ToSQL(dialect)
		if err != nil {
			return "", nil
		}
		lines = append(lines, query)
	}

	return strings.Join(lines, "\n"), nil
}

func (cs *changeSet) CreateTable(table Table) {
	cs.commandBus.Add(newCreateTableCommand(table))
}

func (cs *changeSet) CreateTableWithColumns(tableName string, columns ...*Column) {
	newTable := NewTableWithColumns(tableName, columns...)
	cs.commandBus.Add(newCreateTableCommand(newTable))
}

func (cs *changeSet) CreateTableWithFunc(tableName string, fn CreateTableFunc) {
	// Create a blank table structure that will hold our column definitions.
	newTable := NewTable(tableName)

	// Run the user defined function that will modify the table structure.
	fn(newTable)

	// Create the command.
	command := newCreateTableCommand(newTable)

	// Add the new command to the command bus.
	cs.commandBus.Add(command)
}

func (cs *changeSet) DropTable(name string) {
	cs.commandBus.Add(newDropTableCommand(name))
}

func (cs *changeSet) AlterTable(tableName string, atc *alterTableColumns) {
	cs.commandBus.Add(newAlterTableCommand(tableName, atc))
}

func (cs *changeSet) Execute(backEnd BackEnd) error {
	return cs.commandBus.Execute(backEnd)
}
