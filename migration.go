package evolve

type Migration interface {
	Up(changeSet ChangeSet)
	Down(changeSet ChangeSet)
}

func NewMigrationWrapper(up func(ChangeSet), down func(ChangeSet)) Migration {
	return &migrationWrapper{
		upFunc:   up,
		downFunc: down,
	}
}

type migrationWrapper struct {
	upFunc   func(ChangeSet)
	downFunc func(ChangeSet)
}

func (m *migrationWrapper) Up(changeSet ChangeSet) {
	m.upFunc(changeSet)
}

func (m *migrationWrapper) Down(changeSet ChangeSet) {
	m.downFunc(changeSet)
}
