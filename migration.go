package evolve

type Migration interface {
	Up(changeSet ChangeSet)
	Down(changeSet ChangeSet)
}

func NewMigrationWrapper(up func(ChangeSet), down func(ChangeSet)) Migration {
	return &migration{
		upFunc:   up,
		downFunc: down,
	}
}

type migration struct {
	upFunc   func(ChangeSet)
	downFunc func(ChangeSet)
}

func (m *migration) Up(changeSet ChangeSet) {
	m.upFunc(changeSet)
}

func (m *migration) Down(changeSet ChangeSet) {
	m.downFunc(changeSet)
}
