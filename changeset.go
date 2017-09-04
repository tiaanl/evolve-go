package evolve

type ChangeSet interface {
	CreateTable(tableName string, fn CreateTableFunc)
	DropTable(tableName string)
}

func NewChangeSet(commandBus *commandBus) ChangeSet {
	return &changeSet{
		commandBus: commandBus,
	}
}

type changeSet struct {
	commandBus *commandBus
}

func (cs *changeSet) CreateTable(tableName string, fn CreateTableFunc) {
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
