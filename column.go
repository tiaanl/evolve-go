package evolve

type ColumnType int

const (
	ColumnTypeString ColumnType = iota
	ColumnTypeInteger
	ColumnTypeFloat
	ColumnTypeDateTime
)

type Column struct {
	Name          string
	Type          ColumnType
	Size          int
	AllowNull     bool
	IsPrimary     bool
	AutoIncrement bool
}

func (c *Column) Equals(other *Column) bool {
	return c.Name == other.Name &&
		c.Type == other.Type &&
		c.Size == other.Size &&
		c.AllowNull == other.AllowNull &&
		c.IsPrimary == other.IsPrimary &&
		c.AutoIncrement == other.AutoIncrement
}

func NewColumnPrimary(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeInteger,
		Size:          0,
		AllowNull:     false,
		IsPrimary:     true,
		AutoIncrement: true,
	}
}

func NewColumnString(name string, size int) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeString,
		Size:          size,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func NewColumnInteger(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeInteger,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func NewColumnFloat(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeFloat,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func NewColumnDateTime(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeDateTime,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func newFluentColumn(column *Column) *fluentColumn {
	return &fluentColumn{
		column: column,
	}
}

type fluentColumn struct {
	column *Column
}

func (f *fluentColumn) AllowNull(allowNull bool) *fluentColumn {
	f.column.AllowNull = allowNull
	return f
}

func (f *fluentColumn) IsPrimary(isPrimary bool) *fluentColumn {
	f.column.IsPrimary = isPrimary
	return f
}

func (f *fluentColumn) AutoIncrement(autoIncrement bool) *fluentColumn {
	f.column.AutoIncrement = autoIncrement
	return f
}
