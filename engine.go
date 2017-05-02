package evolve

import (
	"fmt"
)

type Engine interface {
	AddMigration(name string, migration Migration) error

	Up() error
	Down() error
}

func NewEngine(backEnd BackEnd) Engine {
	return &engine{
		backEnd:    backEnd,
		migrations: map[string]Migration{},
		order:      map[int]string{},
		lastIndex:  0,
	}
}

type engine struct {
	backEnd    BackEnd
	migrations map[string]Migration
	order      map[int]string
	lastIndex  int
}

func (e *engine) AddMigration(name string, migration Migration) error {
	// Make sure the key doesn't exist already.
	_, exists := e.migrations[name]
	if exists {
		return fmt.Errorf("Migration with that name already exists (%s)", name)
	}

	// Set the key and migration.
	e.migrations[name] = migration

	e.order[e.lastIndex] = name
	e.lastIndex = e.lastIndex + 1

	return nil
}

func (e *engine) Up() error {
	return e.execute(e.backEnd, up)
}

func (e *engine) Down() error {
	return e.execute(e.backEnd, down)
}

func (e *engine) execute(backEnd BackEnd, fn func(Migration, Schema)) error {
	// Make sure the migrations table exists.
	table := NewTable("migrations")
	table.Primary("id")
	table.String("name", 100)

	err := backEnd.CreateTableIfNotExists(table)
	if err != nil {
		return err
	}

	// Create the command bus we will collect all the migration commands into.
	commandBus := newCommandBus()

	// Create the user friendly schema we'll pass to the user so that they can interact with the command bus.
	schema := NewSchema(commandBus)

	// Run through all the migrations to gather commands into the schema's command bus.
	for _, migrationName := range e.order {
		migration := e.migrations[migrationName]
		fn(migration, schema)
	}

	// Execute all the commands in the command bus and report any errors.
	err = commandBus.Execute(backEnd)
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
