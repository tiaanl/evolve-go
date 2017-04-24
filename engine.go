package evolve

type Engine interface {
	AddMigration(migration Migration)

	Up(backEnd BackEnd) error
	Down(backEnd BackEnd) error
}

func NewEngine() Engine {
	return &engine{
		migrations: []Migration{},
	}
}

type engine struct {
	migrations []Migration
}

func (e *engine) AddMigration(migration Migration) {
	e.migrations = append(e.migrations, migration)
}

func (e *engine) Up(backEnd BackEnd) error {
	return e.execute(backEnd, up)
}

func (e *engine) Down(backEnd BackEnd) error {
	return e.execute(backEnd, down)
}

func (e *engine) execute(backEnd BackEnd, fn func(Migration, Schema)) error {
	// Create the command bus we will collect all the migration commands into.
	commandBus := NewCommandBus()

	// Create the user friendly schema we'll pass to the user so that they can interact with the command bus.
	schema := NewSchema(commandBus)

	// Run through all the migrations to gather commands into the schema's command bus.
	for _, migration := range e.migrations {
		fn(migration, schema)
	}

	// Execute all the commands in the command bus and report any errors.
	err := commandBus.Execute(backEnd)
	if err != nil {
		return err
	}

	return nil
}

func up(m Migration, s Schema) {
	m.Up(s)
}

func down(m Migration, s Schema) {
	m.Down(s)
}
