package evolve

type Migration interface {
	Up(schema Schema)
	Down(schema Schema)
}

func NewMigrationWrapper(up func(Schema), down func(Schema)) Migration {
	return &migrationWrapper{
		upFunc: up,
		downFunc: down,
	}
}

type migrationWrapper struct{
	upFunc func(Schema)
	downFunc func(Schema)
}

func (m *migrationWrapper) Up(schema Schema) {
	m.upFunc(schema)
}

func (m *migrationWrapper) Down(schema Schema) {
	m.downFunc(schema)
}
