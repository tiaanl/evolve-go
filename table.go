package evolve

type Table interface {
	Name() string

	Columns() []*Column
	Column(name string) *Column
	AddColumn(column *Column)

	Primary(name string) *fluentColumn
	String(name string, size int) *fluentColumn
	Integer(name string) *fluentColumn
	DateTime(name string) *fluentColumn
}

func NewTable(name string) Table {
	return &table{
		name:    name,
		columns: []*Column{},
	}
}

func NewTableWithColumns(tableName string, columns []*Column) Table {
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

func (t *table) AddColumn(column *Column) {
	t.columns = append(t.columns, column)
}

func (t *table) Primary(name string) *fluentColumn {
	column := &Column{
		Name:          name,
		Type:          COLUMN_TYPE_UNSIGNED_INTEGER,
		Size:          0,
		AllowNull:     false,
		IsPrimary:     true,
		AutoIncrement: true,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}

func (t *table) String(name string, size int) *fluentColumn {
	column := &Column{
		Name:          name,
		Type:          COLUMN_TYPE_STRING,
		Size:          size,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}

func (t *table) Integer(name string) *fluentColumn {
	column := &Column{
		Name:          name,
		Type:          COLUMN_TYPE_INTEGER,
		Size:          9,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}

func (t *table) DateTime(name string) *fluentColumn {
	column := &Column{
		Name:          name,
		Type:          COLUMN_TYPE_DATE_TIME,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}
