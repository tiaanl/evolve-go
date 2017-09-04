package evolve

type CreateTableFunc func(t Table)

// Schema is the user friendly interface to put Commands into a command bus.
type Schema interface {
	Table(tableName string, fn CreateTableFunc)
	Tables() []Table
}

func NewSchema() Schema {
	return &schema{
		tables: []Table{},
	}
}

type schema struct {
	tables []Table
}

func (s *schema) Tables() []Table {
	return s.tables
}

func (s *schema) Table(tableName string, fn CreateTableFunc) {
	newTable := NewTable(tableName)
	fn(newTable)
	s.tables = append(s.tables, newTable)
}
