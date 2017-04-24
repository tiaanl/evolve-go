package evolve

type Table interface {
	Name() string
	Columns() []Column

	Primary(name string)
	String(name string, size int, allowNull bool)
	DateTime(name string, allowNull bool)
}

func NewTable(name string) Table {
	return &table{
		name: name,
	}
}

type table struct {
	name    string
	columns []Column
}

func (t *table) Name() string {
	return t.name
}

func (t *table) Columns() []Column {
	return t.columns
}

func (t *table) Primary(name string) {
	t.columns = append(t.columns, Column{
		Name:      name,
		Type:      COLUMN_TYPE_UNSIGNED_INTEGER,
		AllowNull: false,
		IsPrimary: true,
	})
}

func (t *table) String(name string, size int, allowNull bool) {
	t.columns = append(t.columns, Column{
		Name:      name,
		Type:      COLUMN_TYPE_STRING,
		Size:      size,
		AllowNull: allowNull,
		IsPrimary: false,
	})
}

func (t *table) DateTime(name string, allowNull bool) {
	t.columns = append(t.columns, Column{
		Name:      name,
		Type:      COLUMN_TYPE_DATE_TIME,
		AllowNull: allowNull,
	})
}
