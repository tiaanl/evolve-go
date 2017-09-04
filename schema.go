package evolve

type CreateTableFunc func(t Table)

// Schema is the user friendly interface to put Commands into a command bus.
type Schema interface {
	Table(tableName string, fn CreateTableFunc)
	Tables() []Table

	CreateTable(tableName string, fn CreateTableFunc)
	DropTable(tableName string)
}

func NewSchema(commandBus *commandBus) Schema {
	return &schema{
		commandBus: commandBus,
		tables:     []Table{},
	}
}

type schema struct {
	commandBus *commandBus
	tables     []Table
}

func (s *schema) Tables() []Table {
	return s.tables
}

func (s *schema) Table(tableName string, fn CreateTableFunc) {
	newTable := NewTable(tableName)
	fn(newTable)
	s.tables = append(s.tables, newTable)
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
