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

func (c *Column) WithName(name string) *Column {
	c.Name = name
	return c
}

func (c *Column) WithType(t ColumnType) *Column {
	c.Type = t
	return c
}

func (c *Column) WithSize(size int) *Column {
	c.Size = size
	return c
}

func (c *Column) Null() *Column {
	c.AllowNull = true
	return c
}

func (c *Column) NotNull() *Column {
	c.AllowNull = false
	return c
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

func StringColumn(name string, size int) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeString,
		Size:          size,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func IntegerColumn(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeInteger,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func FloatColumn(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeFloat,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}

func DateTimeColumn(name string) *Column {
	return &Column{
		Name:          name,
		Type:          ColumnTypeDateTime,
		Size:          0,
		AllowNull:     true,
		IsPrimary:     false,
		AutoIncrement: false,
	}
}
