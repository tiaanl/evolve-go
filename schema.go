package evolve

type CreateTableFunc func(t Table)

// Schema is the user friendly interface to put Commands into a command bus.
type Schema interface {
	CreateTableWithColumns(tableName string, columns []*Column)
	CreateTableWithFunc(tableName string, fn CreateTableFunc)

	GetTable(tableName string) Table
	Tables() []Table
}

func NewSchema() Schema {
	return &schema{
		tables: []Table{},
	}
}

func NewSchemaWithTables(tables []Table) Schema {
	return &schema{
		tables: tables,
	}
}

type schema struct {
	tables []Table
}

func (s *schema) Tables() []Table {
	return s.tables
}

func (s *schema) GetTable(tableName string) Table {
	for _, table := range s.tables {
		if table.Name() == tableName {
			return table
		}
	}

	return nil
}

func (s *schema) CreateTableWithColumns(tableName string, columns []*Column) {
	newTable := NewTableWithColumns(tableName, columns)
	s.tables = append(s.tables, newTable)
}

func (s *schema) CreateTableWithFunc(tableName string, fn CreateTableFunc) {
	newTable := NewTable(tableName)
	fn(newTable)
	s.tables = append(s.tables, newTable)
}
