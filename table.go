package evolve

type Table interface {
	Name() string

	Columns() []*Column
	Column(name string) *Column
	AddColumns(columns ...*Column)
}

func NewTable(name string) Table {
	return &table{
		name:    name,
		columns: []*Column{},
	}
}

func NewTableWithColumns(tableName string, columns ...*Column) Table {
	return &table{
		name:    tableName,
		columns: columns,
	}
}

type table struct {
	name    string
	columns []*Column
}

func (t *table) Name() string {
	return t.name
}

func (t *table) Columns() []*Column {
	return t.columns
}

// Return the column with the specified name.  Returns nil if the column isn't found.
func (t *table) Column(name string) *Column {
	for _, column := range t.columns {
		if column.Name == name {
			return column
		}
	}

	return nil
}

func (t *table) AddColumns(columns ...*Column) {
	for _, column := range columns {
		t.columns = append(t.columns, column)
	}
}
