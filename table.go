package evolve

type Table interface {
	Name() string
	Columns() []*Column

	Primary(name string) *fluentColumn
	String(name string, size int) *fluentColumn
	DateTime(name string) *fluentColumn
}

func NewTable(name string) Table {
	return &table{
		name: name,
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

func (t *table) Primary(name string) *fluentColumn {
	column := &Column{
		Name:      name,
		Type:      COLUMN_TYPE_UNSIGNED_INTEGER,
		Size:      0,
		AllowNull: false,
		IsPrimary: true,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}

func (t *table) String(name string, size int) *fluentColumn {
	column := &Column{
		Name:      name,
		Type:      COLUMN_TYPE_STRING,
		Size:      size,
		AllowNull: true,
		IsPrimary: false,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}

func (t *table) DateTime(name string) *fluentColumn {
	column := &Column{
		Name:      name,
		Type:      COLUMN_TYPE_DATE_TIME,
		Size:      0,
		AllowNull: true,
		IsPrimary: false,
	}

	t.columns = append(t.columns, column)

	return newFluentColumn(column)
}
