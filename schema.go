package evolve

type CreateTableFunc func(t Table)

// Schema is the user friendly interface to put Commands into a command bus.
type Schema interface {
	CreateTable(tableName string, fn CreateTableFunc)
	DropTable(tableName string)
}

func NewSchema(commandBus CommandBus) Schema {
	return &schema{
		commandBus: commandBus,
	}
}

type schema struct {
	commandBus CommandBus
}

func (s *schema) CreateTable(tableName string, fn CreateTableFunc) {
	// Create a blank table structure that will hold our column definitions.
	newTable := NewTable(tableName)

	// Run the user defined function that will modify the table structure.
	fn(newTable)

	// Create the command.
	command := newCreateTableCommand(newTable)

	// Add the new command to the command bus.
	s.commandBus.Add(command)
}

func (s *schema) DropTable(name string) {
	s.commandBus.Add(newDropTableCommand(name))
}
