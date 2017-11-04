package main

import (
	"fmt"

	ev "github.com/tiaanl/evolve-go"
)

func main() {
	table := ev.NewTableWithColumns("users",
		ev.StringColumn("name", 150).NotNull(),
		ev.IntegerColumn("age").Null(),
	)

	dialect := ev.NewDialectMysql()
	sql, _ := dialect.GetCreateTableSQL(table)
	fmt.Println(sql)
}
