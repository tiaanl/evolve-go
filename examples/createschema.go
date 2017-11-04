package main

import (
	"fmt"

	ev "github.com/tiaanl/evolve-go"
)

func main() {
	s := ev.NewSchema()

	s.CreateTableWithColumns("users",
		ev.StringColumn("name", 150).NotNull(),
		ev.IntegerColumn("age").Null(),
	)

	dialect := ev.NewDialectMysql()
	sql, _ := s.ToSQL(dialect)
	fmt.Println(sql)
}
