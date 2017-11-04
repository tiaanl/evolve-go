package main

import (
	"fmt"

	ev "github.com/tiaanl/evolve-go"
)

func main() {
	cs := ev.NewChangeSet()

	cs.CreateTable(ev.NewTableWithColumns("users",
		ev.StringColumn("name", 150).NotNull(),
		ev.IntegerColumn("age").Null(),
	))

	cs.DropTable("users")

	sql, _ := cs.ToSQL(ev.NewDialectMysql())
	fmt.Println(sql)
}
