package evolve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	backEnd := NewBackEndMysql(nil)

	engine := NewEngine()

	engine.AddMigration(NewMigrationWrapper(
		func(schema Schema) {
			schema.CreateTable("users", func(table Table) {
				table.Primary("id")
				table.String("name", 100).AllowNull(false)
				table.String("email", 150).AllowNull(false)
				table.String("password", 100).AllowNull(false)
				table.DateTime("created_at").AllowNull(true)
				table.DateTime("updated_at").AllowNull(true)
			})
		},
		func(schema Schema) {
			schema.DropTable("users")
		},
	))

	err := engine.Up(backEnd)
	assert.NoError(t, err)

	err = engine.Down(backEnd)
	assert.NoError(t, err)
}
