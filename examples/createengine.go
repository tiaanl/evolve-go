package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	ev "github.com/tiaanl/evolve-go"
)

func main() {
	db, _, _ := sqlmock.New()
	engine := ev.NewEngine(ev.NewBackEndMysql(db))

	engine.NewMigration("migration1",
		func(cs ev.ChangeSet) {
			cs.CreateTableWithColumns("users",
				ev.StringColumn("name", 150).NotNull(),
				ev.IntegerColumn("age").Null(),
			)
		},
		func(cs ev.ChangeSet) {
			cs.DropTable("users")
		},
	)

	engine.Update()
}
