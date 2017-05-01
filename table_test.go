package evolve

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTable_Primary(t *testing.T) {
	table := NewTable("atable")

	table.Primary("id")

	assert.Equal(t, "id", table.Column("id").Name)
	assert.Equal(t, COLUMN_TYPE_UNSIGNED_INTEGER, table.Column("id").Type)
	assert.Equal(t, 0, table.Column("id").Size)
	assert.Equal(t, false, table.Column("id").AllowNull)
	assert.Equal(t, true, table.Column("id").IsPrimary)
}

func TestTable_String(t *testing.T) {
	table := NewTable("atable")

	table.String("name", 100)
	table.String("name2", 100).AllowNull(false)

	assert.Equal(t, "name", table.Column("name").Name)
	assert.Equal(t, COLUMN_TYPE_STRING, table.Column("name").Type)
	assert.Equal(t, 100, table.Column("name").Size)
	assert.Equal(t, true, table.Column("name").AllowNull)
	assert.Equal(t, false, table.Column("name").IsPrimary)

	assert.Equal(t, "name2", table.Column("name2").Name)
	assert.Equal(t, COLUMN_TYPE_STRING, table.Column("name2").Type)
	assert.Equal(t, 100, table.Column("name2").Size)
	assert.Equal(t, false, table.Column("name2").AllowNull)
	assert.Equal(t, false, table.Column("name2").IsPrimary)
}
