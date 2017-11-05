package evolve

import (
	"fmt"
)

type Engine interface {
	NewMigration(name string, up func(ChangeSet), down func(ChangeSet))
	AddMigration(name string, migration Migration) error

	Update() error
}

func NewEngine(backEnd BackEnd) Engine {
	return &engine{
		backEnd:       backEnd,
		migrations:    map[string]Migration{},
		order:         map[int]string{},
		lastIndex:     0,
		migrationList: NewMigrationList(backEnd),
	}
}

type engine struct {
	backEnd       BackEnd
	migrations    map[string]Migration
	order         map[int]string
	lastIndex     int
	migrationList MigrationList
}

func (e *engine) NewMigration(name string, up func(ChangeSet), down func(ChangeSet)) {
	e.migrations[name] = &migration{
		upFunc:   up,
		downFunc: down,
	}
}

func (e *engine) AddMigration(name string, migration Migration) error {
	// Make sure the key doesn't exist already.
	_, exists := e.migrations[name]
	if exists {
		return fmt.Errorf("migration with that name already exists (%s)", name)
	}

	// Set the key and migration.
	e.migrations[name] = migration

	e.order[e.lastIndex] = name
	e.lastIndex = e.lastIndex + 1

	return nil
}

func (e *engine) Update() error {
	// Create the user friendly schema we'll pass to the user so that they can interact with the command bus.
	changeSet := NewChangeSet()

	// Run through all the existingMigrations to gather commands into the schema's command bus.
	for _, migrationName := range e.order {
		exists, err := e.migrationList.Exists(migrationName)
		if err != nil {
			return err
		}

		if !exists {
			migration := e.migrations[migrationName]
			migration.Up(changeSet)
			e.migrationList.Add(migrationName)
		}
	}

	// Execute all the commands in the command bus and report any errors.
	err := changeSet.Execute(e.backEnd)
	if err != nil {
		return err
	}

	return nil
}
